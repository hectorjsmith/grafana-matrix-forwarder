package matrix

import (
	"grafana-matrix-forwarder/grafana"
	"testing"
)

func Test_buildFormattedMessageBodyFromAlert(t *testing.T) {
	type args struct {
		alert grafana.AlertPayload
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "alertingStateTest",
			args: args{grafana.AlertPayload{
				State:    "alerting",
				RuleURL:  "http://example.com",
				RuleName: "sample",
				Message:  "sample message",
			}},
			want: "💔 ️<b>ALERT</b><p>Rule: <a href=\"http://example.com\">sample</a> | sample message</p>",
		},
		{
			name: "okStateTest",
			args: args{grafana.AlertPayload{
				State:    "ok",
				RuleURL:  "http://example.com",
				RuleName: "sample",
				Message:  "sample message",
			}},
			want: "💚 ️<b>RESOLVED</b><p>Rule: <a href=\"http://example.com\">sample</a> | sample message</p>",
		},
		{
			name: "noDataStateTest",
			args: args{grafana.AlertPayload{
				State:    "no_data",
				RuleURL:  "http://example.com",
				RuleName: "sample",
				Message:  "sample message",
			}},
			want: "❓️<b>NO DATA</b><ul><p>Rule: <a href=\"http://example.com\">sample</a> | sample message</p>",
		},
		{
			name: "unknownStateTest",
			args: args{grafana.AlertPayload{
				State:    "invalid state",
				RuleURL:  "http://example.com",
				RuleName: "sample",
				Message:  "sample message",
			}},
			want: "❓️<b>UNKNOWN</b><ul><li>Rule: <a href=\"http://example.com\">sample</a> | sample message</li><li>State: <b>invalid state</b></li></ul>",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildFormattedMessageBodyFromAlert(tt.args.alert); got != tt.want {
				t.Errorf("buildFormattedMessageBodyFromAlert() = %v, want %v", got, tt.want)
			}
		})
	}
}
