package main

import (
	"github.com/aws/aws-lambda-go/events"
	"testing"
)

func Test_handle(t *testing.T) {
	type args struct {
		event events.CloudwatchLogsEvent
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				event: events.CloudwatchLogsEvent{
					AWSLogs: events.CloudwatchLogsRawData{
						Data: "H4sIAAAAAAAAAL1STYvbMBD9K0b00EIcybIt2b4Fmm4PLRSSU9ehyJLsFdhWkOQsIeS/d+wkhfbWLfSij6eZN2/m6YIG7b3o9P581KhCHzf7zY+v291u87RFK2RfR+0Aznie8JIkKUlygHvbPTk7HeEFi1ePezE0SuCjs2qSwcfiaOLOxkqf5uMtYRecFgNkUEIJTggmHD+/+7LZb3f7Q0N1kZFStUCTtTIvJROUMFZyxmVCFVD4qfHSmWMwdvxk+qCdR9Uz+q0mlPFxcKbrQPVhKbs96THMkRdkFFRPc05ZXnAgLzilBc2Sgqc8LWgJO0sZyEgymqVplsJKeEpoQUoGCoKBWQUxQNsJgybSPOekzNnqMUOgv9So1yfd16iqkXbOuhqtahT8AsytxwmJCd+TpMpoRcvvy/vguyXg0c1nMapeu3Wnw7cb9P5DFEcLYQRgMGMX3YMXAmfUQkAIK1vSpDFr2izOdCljoTISZ4w1DUyVNg1ZEu7aIAW/2EFjaZzstTS4s9g7iTsTXqZmLe2AhQ/OSqt0bOyfHmNlB2HGB4y9dicj9bqzFWWROo9isFU02sgEPUStnUbQea1HdF39myH8LYaYsbV/5YfT/mhHr982YwfeVCkcQGSY5pI5IXBtrDovTJe6rtE82Xm/PdYPa2ZkXv67PfWsHZb7DH8JkXbqFcSG+QM+Pl/UnCOz5Fzvvh6uPwGVXxwVVAQAAA==",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := handle(tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handle() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
