---
title: Initial Setup
weight: 20
---

## Step 1

Run the forwarder by providing a matrix account to send messages from.

```
$ ./grafana-matrix-forwarder --user @userId:matrix.org --password xxx --homeserver matrix.org
```

## Step 2

Add a new **POST webhook** alert channel with the following target URL: `http://<ip address>:6000/api/v0/forward?roomId=<roomId>`

*Replace with the actual server IP and matrix room ID.*

See the page on [configuration]({{< ref "configuration.md" >}}) to change the port the forwarder listens on.

{{< hint info >}}
NOTE: multiple `roomId` parameters can be provided to forward the alert to multiple rooms at once
{{< /hint >}}

![screenshot of grafana channel setup](img/grafanaChannelSetup.png)

## Step 3

Setup alerts in grafana that are sent to the new alert channel.

![screenshot of grafana alert setup](img/grafanaAlertSetup.png)
