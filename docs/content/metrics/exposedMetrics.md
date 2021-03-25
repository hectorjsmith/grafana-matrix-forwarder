---
title: Exposed Metrics
weight: 10
---

This tool exports prometheus-compatible metrics on the same port at the `/metrics` URL.

{{< hint info >}}
NOTE: All metric names include the `gmf_` prefix (grafana matrix forwarder) to make sure they are unique and make them easier to find.
{{< /hint >}}

## Exposed Metrics
* `up` - Returns 1 if the service is up
* Forward counts
    * `total` - total number of alerts forwarded
    * `success` - number of alerts successfully forwarded
    * `error` - number of alerts where the forwarding process failed (check logs for error details)
* Alert counts by state
    * `total` - total number of alerts received
    * `alerting` - alert count in the *alerting* state
    * `no_data` - alert count in the *no_data* state
    * `ok` - alert count in the *ok* state (resolved alerts)
    * `other` - number of received alerts that have an unknown state (check logs for details)

### Grafana
These metrics can be loaded into Grafana to create a dashboard and/or other alerts. This makes it possible to monitor the forwarder and send an alert of a different type if it goes down (e.g. email).

![Screenshot of an example Grafana dashboard](/img/exampleGrafanaDashboard.png)

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
