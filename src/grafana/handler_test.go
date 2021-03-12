package grafana

import (
	"testing"
)

func Test_buildFormattedMessageBodyFromAlert(t *testing.T) {
	type args struct {
		alert AlertPayload
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "alertingStateTest",
			args: args{AlertPayload{
				State:    "alerting",
				RuleURL:  "http://example.com",
				RuleName: "sample",
				Message:  "sample message",
			}},
			want: "💔 <b>ALERT</b><p>Rule: <a href=\"http://example.com\">sample</a> | sample message</p>",
		},
		{
			name: "alertingStateWithEvalMatchesTest",
			args: args{AlertPayload{
				State:    "alerting",
				RuleURL:  "http://example.com",
				RuleName: "sample",
				Message:  "sample message",
				EvalMatches: []struct {
					Value  float64           `json:"value"`
					Metric string            `json:"metric"`
					Tags   map[string]string `json:"tags"`
				}{
					{
						Value:  10.0,
						Metric: "sample",
						Tags:   map[string]string{},
					},
				},
			}},
			want: "💔 <b>ALERT</b><p>Rule: <a href=\"http://example.com\">sample</a> | sample message</p><ul><li><b>sample</b>: 10</li></ul>",
		},
		{
			name: "okStateTest",
			args: args{AlertPayload{
				State:    "ok",
				RuleURL:  "http://example.com",
				RuleName: "sample",
				Message:  "sample message",
			}},
			want: "💚 <b>RESOLVED</b><p>Rule: <a href=\"http://example.com\">sample</a> | sample message</p>",
		},
		{
			name: "noDataStateTest",
			args: args{AlertPayload{
				State:    "no_data",
				RuleURL:  "http://example.com",
				RuleName: "sample",
				Message:  "sample message",
			}},
			want: "❓ <b>NO DATA</b><p>Rule: <a href=\"http://example.com\">sample</a> | sample message</p>",
		},
		{
			name: "unknownStateTest",
			args: args{AlertPayload{
				State:    "invalid state",
				RuleURL:  "http://example.com",
				RuleName: "sample",
				Message:  "sample message",
			}},
			want: "❓ <b>UNKNOWN</b><p>Rule: <a href=\"http://example.com\">sample</a> | sample message</p>",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := buildFormattedMessageBodyFromAlert(tt.args.alert)
			if err != nil {
				t.Errorf("buildFormattedMessageBodyFromAlert() error: %v", err)
			}
			if got != tt.want {
				t.Errorf("buildFormattedMessageBodyFromAlert() = %v, want %v", got, tt.want)
			}
		})
	}
}
