package grafana

import "fmt"

const (
	// AlertStateAlerting represents the state name grafana uses for alerts that are firing
	AlertStateAlerting = "alerting"
	// AlertStateResolved represents the state name grafana uses for alerts that have been resolved
	AlertStateResolved = "ok"
	// AlertStateNoData represents the state name grafana uses for alerts that are firing because of missing data
	AlertStateNoData = "no_data"
)

// AlertPayload stores the request data sent with the grafana alert webhook
type AlertPayload struct {
	Title       string `json:"title"`
	Message     string `json:"message"`
	RuleName    string `json:"ruleName"`
	RuleURL     string `json:"ruleUrl"`
	State       string `json:"state"`
	OrgID       int    `json:"orgId"`
	DashboardID int    `json:"dashboardId"`
	PanelID     int    `json:"panelId"`
	RuleID      int    `json:"ruleId"`
}

// FullRuleID is defined as the combination of the OrgID, DashboardID, PanelID, and RuleID
func (payload AlertPayload) FullRuleID() string {
	return fmt.Sprintf("%d.%d.%d.%d", payload.OrgID, payload.DashboardID, payload.PanelID, payload.RuleID)
}
