package server

import (
	"bytes"
	"html/template"
)

type metricData struct {
	MetricName   string
	MetricType   string
	MetricLabels map[string]map[string]float32
	MetricValue  float32
}

const (
	metricNamePrefix  = "gmf_"
	metricTemplateStr = `# HELP {{ .MetricName }}
#TYPE {{ .MetricName }} {{ .MetricType }}
{{- $labelLen := len .MetricLabels -}}
{{- if eq 0 $labelLen }}
{{ .MetricName }} {{ .MetricValue }}
{{ else }}
{{- range $labelName, $labelValueMap := .MetricLabels -}}
{{- range $labelValue, $value := $labelValueMap }}
{{ $.MetricName }}{{ "{" }}"{{ $labelName }}"="{{ $labelValue }}"{{ "}" }} {{ $value }}
{{- end -}}
{{- end -}}
{{- end }}
`
)

var metricTemplate = template.Must(template.New("metric").Parse(metricTemplateStr))

func (serverMetrics serverMetrics) buildMetrics() (metrics string, err error) {
	var buffer string
	buffer, err = buildMetricDataString(buildSimpleMetricData("up", "gauge", 1.0))
	if err != nil {
		return
	}
	metrics += buffer

	buffer, err = buildMetricDataString(buildMetricDataWithLabel(
		"forwards",
		"gauge",
		"result",
		[]string{
			"total",
			"success",
			"error"},
		[]float32{
			float32(serverMetrics.totalForwardCount),
			float32(serverMetrics.successForwardCount),
			float32(serverMetrics.failForwardCount)}))

	if err != nil {
		return
	}
	metrics += buffer
	return
}

func buildMetricDataString(metricData metricData) (string, error) {
	buffer := new(bytes.Buffer)
	err := metricTemplate.Execute(buffer, metricData)
	return buffer.String(), err
}

func buildSimpleMetricData(metricName, metricType string, metricValue float32) metricData {
	return metricData{
		MetricName:  metricNamePrefix + metricName,
		MetricType:  metricType,
		MetricValue: metricValue,
	}
}

func buildMetricDataWithLabel(metricName, metricType, labelTypeName string, labels []string, values []float32) metricData {
	labelData := map[string]map[string]float32{}
	labelTypeData := map[string]float32{}
	for i := 0; i < len(labels); i++ {
		labelTypeData[labels[i]] = values[i]
	}
	labelData[labelTypeName] = labelTypeData
	return metricData{
		MetricName:   metricNamePrefix + metricName,
		MetricType:   metricType,
		MetricLabels: labelData,
	}
}
