package formatter

import (
	"fmt"
	htmlTemplate "html/template"
	textTemplate "text/template"
)

const (
	alertMessageTemplateStr = `{{ .StateEmoji }} <b>{{ .StateStr }}</b>
{{- with .Payload }}<p>Rule: <a href="{{ .RuleURL }}">{{ .RuleName }}</a>{{ if .Message }} | {{ .Message }}{{ end }}</p>
{{- if gt (len .EvalMatches) 0 }}<ul>{{ range $match := .EvalMatches }}<li><b>{{ .Metric }}</b>: {{ RoundValue .Value $.MetricRounding }}</li>{{ end }}</ul>{{ end }}
{{- if gt (len .Tags) 0 }}<p>Tags:</p><ul>{{ range $tagKey, $tagValue := .Tags }}<li><b>{{ $tagKey }}</b>: {{ $tagValue }}</li>{{ end }}</ul>{{ end }}{{ end }}`
	resolvedReactionStr  = `✅`
	resolveReplyStr      = "<mx-reply><blockquote>{{ . }}</blockquote></mx-reply>💚 ️<b>RESOLVED</b>"
	resolveReplyPlainStr = `💚 ️RESOLVED`
)

var (
	alertMessageTemplate = htmlTemplate.Must(htmlTemplate.New("alertMessage").Funcs(htmlTemplate.FuncMap{
		"RoundValue": roundMetricValue,
	}).Parse(alertMessageTemplateStr))
	resolveReplyTemplate = textTemplate.Must(textTemplate.New("resolveReply").Parse(resolveReplyStr))
)

func roundMetricValue(rawValue float64, metricRounding int) string {
	var format string
	if metricRounding >= 0 {
		format = fmt.Sprintf("%%.%df", metricRounding)
	} else {
		format = "%v"
	}
	return fmt.Sprintf(format, rawValue)
}
