package v0

import (
	"fmt"
	"grafana-matrix-forwarder/forwarder"
)

// AlertPayload stores the request data sent with the grafana alert webhook
type AlertPayload struct {
	Title       string `json:"title"`
	Message     string `json:"message"`
	State       string `json:"state"`
	RuleName    string `json:"ruleName"`
	RuleURL     string `json:"ruleUrl"`
	RuleID      int    `json:"ruleId"`
	OrgID       int    `json:"orgId"`
	DashboardID int    `json:"dashboardId"`
	PanelID     int    `json:"panelId"`
	EvalMatches []struct {
		Value  float64           `json:"value"`
		Metric string            `json:"metric"`
		Tags   map[string]string `json:"tags"`
	} `json:"evalMatches"`
	Tags map[string]string `json:"tags"`
}

// FullRuleID is defined as the combination of the OrgID, DashboardID, PanelID, and RuleID
func (payload AlertPayload) FullRuleID() string {
	return fmt.Sprintf("%d.%d.%d.%d", payload.OrgID, payload.DashboardID, payload.PanelID, payload.RuleID)
}

func (payload AlertPayload) ToForwarderData() forwarder.Data {
	return forwarder.Data{
		Id:       payload.FullRuleID(),
		State:    payload.State,
		RuleURL:  payload.RuleURL,
		RuleName: payload.RuleName,
		Message:  payload.Message,
		Tags:     payload.Tags,
		EvalMatches: []struct {
			Value  float64
			Metric string
			Tags   map[string]string
		}(payload.EvalMatches),
	}
}
