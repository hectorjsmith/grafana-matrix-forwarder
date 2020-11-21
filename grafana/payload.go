package grafana

type AlertPayload struct {
	Title    string `json:"title"`
	Message  string `json:"message"`
	RuleName string `json:"ruleName"`
	RuleUrl  string `json:"ruleUrl"`
	State    string `json:"state"`
}
