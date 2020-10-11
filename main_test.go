package main

import (
	"github.com/aws/aws-lambda-go/events"
	"testing"
)

func Test_handle(t *testing.T) {
	host = "listener.logz.io"
	port = "8071"
	token = "palOPqZOzEnDKhrVYZvgEmZzJqWRaVdp"

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
						Data: "H4sIAAAAAAAAAK1TW0/bMBT+K5G1x4b4+O68FdEhpDEQ7fYwgiY3dkqktukSt4wh/vtOUjaxCxKMPcX5jn2+y7HvyCp0nVuE2e0mkJwcjWfjz6eT6XR8PCEj0tysQ4uw0BK0pcApSISXzeK4bbYbrGTupsuWbjX3Ltu0jd+WsUvdpk4XTerDrl/uD0xjG9wKTzDKaAY0ozq7fPNuPJtMZ1fGaKdL6UAKIzQEWyltrRBsHpTzwLFFt513ZVtvYt2s39bLGNqO5JfkF06k6dLY1osFqr4aaCe7sI79zjtSe2TnUjNtlAVLJZNCCKCGg2KSUw6CWVDaMCUF1yAoN8oA11QxgQpijVlFt0LboNCE5VRxruzoR4bYfjobX8ySi/Bli1tPfJ5IrZQvPUuDcTwVUunUUlelc2EN5WjZViL5iGbQVp485FGsyf3oT8EaGVGnpMoaxfBrgWlplRDSYHIWRSstOXCpQD8p2MBjwXcFWYZdWBYkL0i9rpqCjAoSu+G/H1UKNKV6BjqXkFP+aaivusWwod0bHbC29gNmfWlLW81Tx7lIBdNlapE7Nb4SvuIeNNX7JiFeN/sz52dousc2Ll4PSLaDn/epIPd/S4QBZ30eYLVkDIxVmDVOy1rOqTUMV1pZyhQXEsdknkhE4JRfkYj4PZFu06y78G+RtBhlDpQKXKPOuO1ZGaX4O2/87dDs/6Rhn5nG5P3RS6/zq9VpSp+p7mJyfvby91bEo23r4vDigII+UCpZdUU8rJfL4JNHRRzFUDkNq6a9Tab1t4AoM8npIYLua/JQ+NCFnlsN+Mm6jo+bGHlg9m36ZK7uvwOxudMpdAUAAA==",
					},
				},
			},
			wantErr: false,
		},
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
