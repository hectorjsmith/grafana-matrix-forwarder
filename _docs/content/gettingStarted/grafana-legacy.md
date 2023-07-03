---
title: 'Grafana: Legacy'
weight: 50
---

{{< hint warning >}}
This is the old alert system used by Grafana. It has now been deprecated in favour of a [unified]({{< ref "grafana-unified.md" >}}) alert system.
{{< /hint >}}

## Step 1

Add a new **POST webhook** alert channel with the following target URL: `http://<ip address>:6000/api/v1/standard?roomId=<roomId>`

*Replace with the actual server IP and matrix room ID.*

{{< hint info >}}
NOTE: multiple `roomId` parameters can be provided to forward the alert to multiple rooms at once
{{< /hint >}}

![screenshot of grafana channel setup](img/grafanaLegacyChannelSetup.png)

## Step 2

Setup alerts in grafana that are sent to the new alert channel.

![screenshot of grafana alert setup](img/grafanaLegacyAlertSetup.png)
