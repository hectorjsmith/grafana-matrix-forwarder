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
	Annotations  map[string]string `json:"annotations"`
	Labels       map[string]string `json:"labels"`
	DashboardUrl string            `json:"dashboardURL"`
	Fingerprint  string            `json:"fingerprint"`
}

// FullRuleID is defined as the combination of the OrgID, DashboardID, PanelID, and RuleID
func (payload alertPayload) FullRuleID() string {
	return fmt.Sprintf("unified.%d.%s", payload.OrgID, payload.Alerts[0].Fingerprint)
}

func (payload alertPayload) ToForwarderData() model.AlertData {
	return model.AlertData{
		Id:       payload.FullRuleID(),
		State:    payload.State,
		RuleURL:  payload.Alerts[0].DashboardUrl,
		RuleName: payload.CommonLabels["alertname"],
		Message:  payload.CommonAnnotations["description"],
		Tags:     payload.CommonLabels,
		EvalMatches: []struct {
			Value  float64
			Metric string
			Tags   map[string]string
		}{},
	}
}
