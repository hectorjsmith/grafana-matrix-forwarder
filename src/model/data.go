package model

const (
	// AlertStateAlerting represents the state name grafana uses for alerts that are firing
	AlertStateAlerting = "alerting"
	// AlertStateResolved represents the state name grafana uses for alerts that have been resolved
	AlertStateResolved = "ok"
	// AlertStateNoData represents the state name grafana uses for alerts that are firing because of missing data
	AlertStateNoData = "no_data"
)

type AlertData struct {
	Id          string
	State       string
	RuleURL     string
	RuleName    string
	Message     string
	RawData     string
	Tags        map[string]string
	EvalMatches []struct {
		Value  float64
		Metric string
		Tags   map[string]string
	}
}
