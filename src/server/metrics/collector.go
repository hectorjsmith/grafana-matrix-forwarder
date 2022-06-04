package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"grafana-matrix-forwarder/model"
	v0 "grafana-matrix-forwarder/server/v0"
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

func (c *Collector) IncrementSuccess() {
	c.successForwardCount++
}

func (c *Collector) IncrementFailure() {
	c.failForwardCount++
}

func (c *Collector) RecordAlert(alert v0.AlertPayload) {
	if alert.State == model.AlertStateAlerting {
		c.alertingAlertCount++
	} else if alert.State == model.AlertStateResolved {
		c.resolvedAlertCount++
	} else if alert.State == model.AlertStateNoData {
		c.noDataAlertCount++
	} else {
		c.otherAlertCount++
	}
}
