package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"grafana-matrix-forwarder/model"
)

type Collector struct {
	successForwardCount int
	failForwardCount    int
	alertCountByState   map[string]int
}

func NewCollector() *Collector {
	return &Collector{
		successForwardCount: 0,
		failForwardCount:    0,
		alertCountByState:   map[string]int{},
	}
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
	for state, count := range c.alertCountByState {
		ch <- prometheus.MustNewConstMetric(
			metricAlertCount, prometheus.CounterValue, float64(count), state,
		)
	}
}

func (c *Collector) IncrementSuccess() {
	c.successForwardCount++
}

func (c *Collector) IncrementFailure() {
	c.failForwardCount++
}

func (c *Collector) RecordAlert(alert model.AlertData) {
	if count, ok := c.alertCountByState[alert.State]; !ok {
		c.alertCountByState[alert.State] = 1
	} else {
		c.alertCountByState[alert.State] = count + 1
	}
}
