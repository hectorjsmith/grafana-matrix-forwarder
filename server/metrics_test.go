package server

import "testing"

func Test_serverMetrics_buildMetrics(t *testing.T) {
	type fields struct {
		totalForwardCount   int
		successForwardCount int
		failForwardCount    int
	}
	tests := []struct {
		name        string
		fields      fields
		wantMetrics string
		wantErr     bool
	}{
		{
			name: "happyPathTest",
			fields: fields{
				totalForwardCount:   10,
				successForwardCount: 4,
				failForwardCount:    6,
			},
			wantMetrics: `# HELP gmf_up
#TYPE gmf_up gauge
gmf_up 1

# HELP gmf_forwards
#TYPE gmf_forwards gauge
gmf_forwards{"result"="error"} 6
gmf_forwards{"result"="success"} 4
gmf_forwards{"result"="total"} 10
`,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			serverMetrics := serverMetrics{
				totalForwardCount:   tt.fields.totalForwardCount,
				successForwardCount: tt.fields.successForwardCount,
				failForwardCount:    tt.fields.failForwardCount,
			}
			gotMetrics, err := serverMetrics.buildMetrics()
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
