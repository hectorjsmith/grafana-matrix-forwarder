package grafana

type AlertPayload struct {
	Title   string `json:"title"`
	Message string `json:"message"`
	RuleUrl string `json:"ruleUrl"`
}
