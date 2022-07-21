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
}

// FullRuleID is defined as the combination of the OrgID, DashboardID, PanelID, and RuleID
func (payload alertPayload) FullRuleID() string {
	return fmt.Sprintf("unified.%d.%s", payload.OrgID, payload.Alerts[0].Fingerprint)
}

func (payload alertPayload) ToForwarderData() []model.AlertData {
	data := make([]model.AlertData, len(payload.Alerts))
	for i, alert := range payload.Alerts {

		data[i] = model.AlertData{
			Id:       payload.FullRuleID(),
			State:    payload.State,
			RuleURL:  alert.PanelUrl,
			RuleName: alert.Labels["alertname"],
			Message:  alert.Annotations["summary"],
			Tags:     alert.Labels,
			EvalMatches: []struct {
				Value  float64
				Metric string
				Tags   map[string]string
			}{},
		}
	}
	return data
}
