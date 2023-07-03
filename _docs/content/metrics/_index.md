---
title: Metrics
weight: -10
---

This tool exports prometheus-compatible metrics on the same port at the `/metrics` URL.

{{< hint info >}}
NOTE: All metric names include the `gmf_` prefix (grafana matrix forwarder) to make sure they are unique and make them easier to find.
{{< /hint >}}

## Exposed Metrics

| Metric Name | Type | Description |
|--------|------|-------------|
| `gmf_up` | `gauge` | Returns 1 if the service is up |
| `gmf_forwards` | `gauge` | Counts the number of matrix messages sent (both successfully and with errors) |
| `gmf_alerts` | `gauge` | Counts the number of grafana alerts received by type |
