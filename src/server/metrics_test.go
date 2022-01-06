package server

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func Test_serverMetrics_buildMetrics1(t *testing.T) {
	type fields struct {
		totalForwardCount   int
		successForwardCount int
		failForwardCount    int
		alertingAlertCount  int
		resolvedAlertCount  int
		noDataAlertCount    int
		otherAlertCount     int
	}
	tests := []struct {
		name        string
		fields      fields
		wantMetrics string
		wantErr     bool
	}{{
		name: "happyPathTest",
		fields: fields{
			successForwardCount: 4,
			failForwardCount:    6,
			alertingAlertCount:  5,
			resolvedAlertCount:  3,
			noDataAlertCount:    1,
			otherAlertCount:     1,
		},
		wantMetrics: `# HELP gmf_alerts Alert states being processed by the forwarder
# TYPE gmf_alerts counter
gmf_alerts{state="alerting"} 5
gmf_alerts{state="no_data"} 1
gmf_alerts{state="ok"} 3
gmf_alerts{state="other"} 1
# HELP gmf_forwards Successful and failed alert forwards
# TYPE gmf_forwards counter
gmf_forwards{result="fail"} 6
gmf_forwards{result="success"} 4
# HELP gmf_up Alert forwarder is up and running
# TYPE gmf_up gauge
gmf_up 1
`,
		wantErr: false,
	},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			serverMetrics := &Collector{
				successForwardCount: tt.fields.successForwardCount,
				failForwardCount:    tt.fields.failForwardCount,
				alertingAlertCount:  tt.fields.alertingAlertCount,
				resolvedAlertCount:  tt.fields.resolvedAlertCount,
				noDataAlertCount:    tt.fields.noDataAlertCount,
				otherAlertCount:     tt.fields.otherAlertCount,
			}
			registry := prometheus.NewRegistry()
			registry.MustRegister(serverMetrics)

			s := httptest.NewServer(promhttp.HandlerFor(registry, promhttp.HandlerOpts{}))
			defer s.Close()

			resp, err := s.Client().Get(s.URL)
			gotMetricsBytes, err := ioutil.ReadAll(resp.Body)
			gotMetrics := string(gotMetricsBytes)

			if (err != nil) != tt.wantErr {
				t.Errorf("buildMetrics() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotMetrics != tt.wantMetrics {
				t.Errorf("buildMetrics() gotMetrics = %v, want %v", gotMetrics, tt.wantMetrics)
			}
		})
	}
}
