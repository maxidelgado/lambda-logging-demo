package main

import (
	"errors"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/go-resty/resty/v2"
	json "github.com/json-iterator/go"
	"go.uber.org/zap"
)

var (
	host  = os.Getenv("HOST")
	port  = os.Getenv("PORT")
	token = os.Getenv("TOKEN")

	log, _ = zap.NewProduction()
)

func main() {
	lambda.Start(handle)
}

func handle(event events.CloudwatchLogsEvent) error {
	data, err := event.AWSLogs.Parse()
	if err != nil {
		log.Error("processAll()", zap.Error(err))
		return err
	}

	if err := processAll(data.LogGroup, data.LogStream, data.LogEvents); err != nil {
		log.Error("processAll()", zap.Error(err))
		return err
	}

	return nil
}

func processAll(group, stream string, logs []events.CloudwatchLogsLogEvent) error {
	addr := "https://" + host + ":" + port

	client := resty.New()
	req := client.R().
		SetQueryParam("token", token)

	for _, log := range logs {
		raw, _ := logMessage(group, stream, log)
		if raw == nil {
			continue
		}

		_, err := req.SetBody(raw).Post(addr)
		if err != nil {
			return err
		}
	}

	return nil
}

func logMessage(group, stream string, event events.CloudwatchLogsLogEvent) ([]byte, error) {
	if strings.Contains(event.Message, "START RequestId") ||
		strings.Contains(event.Message, "END RequestId") ||
		strings.Contains(event.Message, "REPORT RequestId") {
		return nil, errors.New("skipped log: START - END - REPORT")
	}

	funcName := functionName(group)
	funcVersion := lambdaVersion(stream)

	msg := logMsg{
		Stream:        stream,
		Group:         group,
		LambdaName:    funcName,
		Type:          "cloudwatch",
		Token:         token,
		Message:       []byte(event.Message),
		LambdaVersion: funcVersion,
	}

	raw, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}

	return raw, nil
}

func lambdaVersion(stream string) string {
	start := strings.Index(stream, "[")
	end := strings.Index(stream, "]")

	return stream[start+1 : end]
}

func functionName(group string) string {
	arr := strings.Split(group, "/")
	return arr[len(arr)-1]
}

type logMsg struct {
	Stream        string `json:"stream"`
	Group         string `json:"group"`
	LambdaName    string `json:"lambda_name"`
	Type          string `json:"type"`
	Token         string `json:"token"`
	LambdaVersion string `json:"lambda_version"`
	Message       []byte `json:"message"`
}
