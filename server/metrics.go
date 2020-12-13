package server

import (
	"bytes"
	"html/template"
)

type metricData struct {
	MetricName  string
	MetricType  string
	MetricValue float32
}

const metricNamePrefix = "gmf_"
const metricTemplateStr = "# HELP {{ .MetricName }}\n" +
	"#TYPE {{ .MetricName }} {{ .MetricType }}\n" +
	"{{ .MetricName }} {{ .MetricValue }}\n"

var metricTemplate = template.Must(template.New("metric").Parse(metricTemplateStr))

func buildMetricString(metricName, metricType string, metricValue float32) (string, error) {
	data := metricData{
		MetricName:  metricNamePrefix + metricName,
		MetricType:  metricType,
		MetricValue: metricValue,
	}
	buffer := new(bytes.Buffer)
	err := metricTemplate.Execute(buffer, data)
	return buffer.String(), err
}
