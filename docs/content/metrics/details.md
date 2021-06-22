---
title: Details
weight: 10
---

## `gmf_forwards`

This metric counts the number of alerts which were forwarded on to a matrix room. It splits into three labels:

| Label Name | Description |
|------------|-------------|
| `gmf_forwards{result="total"}` | Total number of matrix messages forwarded |
| `gmf_forwards{result="success"}` | Number of matrix messages which were forwarded successfully |
| `gmf_forwards{result="error"}` | Number of matrix messages where the forwarding process failed  |

## `gmf_alerts`

This metric counts the number of grafana alerts received by the tool. It splits into five labels:

| Label Name | Description |
|------------|-------------|
| `gmf_alerts{state="total"}` | Total number of grafana alerts received |
| `gmf_alerts{state="alerting"}` | Number of grafana alerts with the *alerting* state (i.e. the alert is firing) |
| `gmf_alerts{state="no_data"}` | Number of grafana alerts with the *no_data* state (for example, when grafana failed to read data for that metric) |
| `gmf_alerts{state="ok"}` | Number of grafana alerts with the *ok* state (i.e. resolved alerts) |
| `gmf_alerts{state="other"}` | Number of alerts with an unknown state - check logs for details |

### Raw Output Example
{{< highlight bash "linenos=table" >}}
# HELP gmf_up
# TYPE gmf_up gauge
gmf_up 1
# HELP gmf_forwards
# TYPE gmf_forwards gauge
gmf_forwards{result="error"} 1
gmf_forwards{result="success"} 5
gmf_forwards{result="total"} 6
# HELP gmf_alerts
# TYPE gmf_alerts gauge
gmf_alerts{state="alerting"} 1
gmf_alerts{state="no_data"} 1
gmf_alerts{state="ok"} 2
gmf_alerts{state="other"} 1
gmf_alerts{state="total"} 6
{{< /highlight >}}
