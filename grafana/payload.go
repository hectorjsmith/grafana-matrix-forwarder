package grafana

const (
	AlertStateAlerting = "alerting"
	AlertStateResolved = "ok"
	AlertStateNoData   = "no_data"
)

type AlertPayload struct {
	Title    string `json:"title"`
	Message  string `json:"message"`
	RuleName string `json:"ruleName"`
	RuleUrl  string `json:"ruleUrl"`
	State    string `json:"state"`
}
