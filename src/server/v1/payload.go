package v1

import (
	"fmt"
	"grafana-matrix-forwarder/model"
)

type alertPayload struct {
	Title             string            `json:"title"`
	Message           string            `json:"message"`
	State             string            `json:"state"`
	OrgID             int               `json:"orgId"`
	GroupLabels       map[string]string `json:"groupLabels"`
	CommonLabels      map[string]string `json:"commonLabels"`
	CommonAnnotations map[string]string `json:"commonAnnotations"`
	Alerts            []alert           `json:"alerts"`
}

type alert struct {
	Status       string            `json:"status"`
	Annotations  map[string]string `json:"annotations"`
	Labels       map[string]string `json:"labels"`
	DashboardUrl string            `json:"dashboardURL"`
	PanelUrl     string            `json:"panelURL"`
	Fingerprint  string            `json:"fingerprint"`
	ValueString  string            `json:"valueString"`
}

// FullRuleID is defined as the combination of the OrgID, DashboardID, PanelID, and RuleID
func fullRuleID(p alertPayload, a alert) string {
	return fmt.Sprintf("unified.%d.%s", p.OrgID, a.Fingerprint)
}

func normaliseStatus(status string) string {
	switch status {
	case "firing":
		return model.AlertStateAlerting
	case "resolved":
		return model.AlertStateResolved
	default:
		return status
	}
}

func (payload alertPayload) ToForwarderData() []model.AlertData {
	data := make([]model.AlertData, len(payload.Alerts))
	for i, alert := range payload.Alerts {
		data[i] = model.AlertData{
			Id:       fullRuleID(payload, alert),
			State:    normaliseStatus(alert.Status),
			RuleURL:  alert.PanelUrl,
			RuleName: alert.Labels["alertname"],
			Message:  alert.Annotations["summary"],
			RawData:  alert.ValueString,
			Tags:     map[string]string{},
			EvalMatches: []struct {
				Value  float64
				Metric string
				Tags   map[string]string
			}{},
		}
	}
	return data
}
