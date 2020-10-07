package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/go-resty/resty/v2"
	json "github.com/json-iterator/go"
)

var (
	host  = os.Getenv("HOST")
	port  = os.Getenv("PORT")
	token = os.Getenv("TOKEN")
)

func main() {
	lambda.Start(handle)
}

func handle(event []byte) error {
	var m map[string]interface{}
	_ = json.Unmarshal(event, &m)
	fmt.Printf("event: %v", m)
	/*
		if err := processAll(event.LogGroup, event.LogStream, event.LogEvents); err != nil {
			fmt.Printf("error: %v", err.Error())
			return err
		}
	*/
	return nil
}

func processAll(group, stream string, logs []events.CloudwatchLogsLogEvent) error {
	addr := host + ":" + port

	client := resty.New()
	req := client.R().
		SetQueryParam("token", token)

	for _, log := range logs {
		raw := logMessage(group, stream, log)
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

func logMessage(group, stream string, event events.CloudwatchLogsLogEvent) []byte {
	if strings.Contains(event.Message, "START RequestId") ||
		strings.Contains(event.Message, "END RequestId") ||
		strings.Contains(event.Message, "REPORT RequestId") {
		return nil
	}

	funcName := functionName(group)
	funcVersion, err := lambdaVersion(stream)
	if err != nil {
		return nil
	}

	msg := log{
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
		return nil
	}

	return raw
}

func lambdaVersion(stream string) (int, error) {
	start := strings.Index(stream, "[")
	end := strings.Index(stream, "]")

	return strconv.Atoi(stream[start:end])
}

func functionName(group string) string {
	return strings.Split(group, "/")[len(group)-1]
}

type log struct {
	Stream        string `json:"stream"`
	Group         string `json:"group"`
	LambdaName    string `json:"lambda_name"`
	Type          string `json:"type"`
	Token         string `json:"token"`
	Message       []byte `json:"message"`
	LambdaVersion int    `json:"lambda_version"`
}
