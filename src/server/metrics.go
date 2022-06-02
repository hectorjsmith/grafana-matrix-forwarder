package server

import (
	"github.com/prometheus/client_golang/prometheus"
	"grafana-matrix-forwarder/server/v0"
)

const namespace = "gmf"

var (
	collectorInstance = &Collector{}

	upMetric = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "up"),
		"Alert forwarder is up and running",
		nil, nil,
	)
	metricForwardCount = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "forwards"),
		"Successful and failed alert forwards",
		[]string{"result"}, nil,
	)
	metricAlertCount = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "alerts"),
		"Alert states being processed by the forwarder",
		[]string{"state"}, nil,
	)
)

type Collector struct {
	successForwardCount int
	failForwardCount    int
	alertingAlertCount  int
	resolvedAlertCount  int
	noDataAlertCount    int
	otherAlertCount     int
}

func (c *Collector) Describe(ch chan<- *prometheus.Desc) {
	ch <- metricForwardCount
	ch <- metricAlertCount
	ch <- upMetric
}

func (c *Collector) Collect(ch chan<- prometheus.Metric) {
	ch <- prometheus.MustNewConstMetric(upMetric, prometheus.GaugeValue, float64(1))
	c.collectForwardCount(ch)
	c.collectAlertCount(ch)
}

func (c *Collector) collectForwardCount(ch chan<- prometheus.Metric) {
	ch <- prometheus.MustNewConstMetric(
		metricForwardCount, prometheus.CounterValue, float64(c.successForwardCount), "success",
	)
	ch <- prometheus.MustNewConstMetric(
		metricForwardCount, prometheus.CounterValue, float64(c.failForwardCount), "fail",
	)
}

func (c *Collector) collectAlertCount(ch chan<- prometheus.Metric) {
	ch <- prometheus.MustNewConstMetric(
		metricAlertCount, prometheus.CounterValue, float64(c.alertingAlertCount), "alerting",
	)
	ch <- prometheus.MustNewConstMetric(
		metricAlertCount, prometheus.CounterValue, float64(c.resolvedAlertCount), "ok",
	)
	ch <- prometheus.MustNewConstMetric(
		metricAlertCount, prometheus.CounterValue, float64(c.noDataAlertCount), "no_data",
	)
	ch <- prometheus.MustNewConstMetric(
		metricAlertCount, prometheus.CounterValue, float64(c.otherAlertCount), "other",
	)
}

func updateAlertMetrics(alert v0.AlertPayload) {
	if alert.State == v0.AlertStateAlerting {
		collectorInstance.alertingAlertCount++
	} else if alert.State == v0.AlertStateResolved {
		collectorInstance.resolvedAlertCount++
	} else if alert.State == v0.AlertStateNoData {
		collectorInstance.noDataAlertCount++
	} else {
		collectorInstance.otherAlertCount++
	}
}
