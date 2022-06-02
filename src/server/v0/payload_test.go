package v0

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestAlertPayload(t *testing.T) {
	tests := []struct {
		name     string
		jsonData string
		want     *AlertPayload
	}{
		{
			name: "GIVEN standard grafana alert payload WHEN converted to struct THEN data matches expected",
			jsonData: `{
  "dashboardId": 13,
  "evalMatches": [
    {
      "value": 1.4,
      "metric": "example.host:9100",
      "tags": {
        "__name__": "node_load1",
        "instance": "example.host:9100",
        "job": "node"
      }
    }
  ],
  "message": "This is a sample alert - please ignore",
  "orgId": 1,
  "panelId": 2,
  "ruleId": 26,
  "ruleName": "My Test Alert",
  "ruleUrl": "https://example.com/d/-IaNDf8Gz/testing-dashboard?tab=alert\\u0026viewPanel=2\\u0026orgId=1",
  "state": "alerting",
  "tags": {
    "priority": "high"
  },
  "title": "[Alerting] My Test Alert"
}`,
			want: &AlertPayload{
				Title:       "[Alerting] My Test Alert",
				Message:     "This is a sample alert - please ignore",
				State:       "alerting",
				RuleName:    "My Test Alert",
				RuleURL:     "https://example.com/d/-IaNDf8Gz/testing-dashboard?tab=alert\\u0026viewPanel=2\\u0026orgId=1",
				RuleID:      26,
				OrgID:       1,
				DashboardID: 13,
				PanelID:     2,
				EvalMatches: []struct {
					Value  float64           `json:"value"`
					Metric string            `json:"metric"`
					Tags   map[string]string `json:"tags"`
				}{
					{
						Value:  1.4,
						Metric: "example.host:9100",
						Tags:   map[string]string{"__name__": "node_load1", "instance": "example.host:9100", "job": "node"},
					},
				},
				Tags: map[string]string{"priority": "high"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			payload, err := convertJsonToPayload(tt.jsonData)
			if err != nil {
				t.Errorf("Failed to convert data - %v", err)
			}
			var payloadStr = fmt.Sprintf("%v", *payload)
			var wantStr = fmt.Sprintf("%v", *tt.want)
			if payloadStr != wantStr {
				t.Errorf("got %s, want %s", payloadStr, wantStr)
			}
		})
	}
}

func convertJsonToPayload(jsonData string) (*AlertPayload, error) {
	var alertPayload *AlertPayload
	err := json.Unmarshal([]byte(jsonData), &alertPayload)
	return alertPayload, err
}
