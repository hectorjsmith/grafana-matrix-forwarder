package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

const namespace = "gmf"

var (
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
