---
title: Details
weight: 10
---

## `gmf_forwards`

This metric counts the number of alerts which were forwarded on to a matrix room. It splits into three labels:

| Label Name | Description |
|------------|-------------|
| `gmf_forwards{result="success"}` | Number of matrix messages which were forwarded successfully |
| `gmf_forwards{result="error"}` | Number of matrix messages where the forwarding process failed  |

## `gmf_alerts`

This metric counts the number of grafana alerts received by the tool. It splits into five labels:

| Label Name | Description |
|------------|-------------|
| `gmf_alerts{state="alerting"}` | Number of grafana alerts with the *alerting* state (i.e. the alert is firing) |
| `gmf_alerts{state="no_data"}` | Number of grafana alerts with the *no_data* state (for example, when grafana failed to read data for that metric) |
| `gmf_alerts{state="ok"}` | Number of grafana alerts with the *ok* state (i.e. resolved alerts) |
| `gmf_alerts{state="other"}` | Number of alerts with an unknown state - check logs for details |

### Raw Output Example
{{< highlight bash "linenos=table" >}}
# HELP gmf_alerts Alert states being processed by the forwarder
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
{{< /highlight >}}
