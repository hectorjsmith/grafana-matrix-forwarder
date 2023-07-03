---
title: 'Grafana: Unified'
weight: 50
---

## Step 1

Add a new **Contact Point** of type **Webhook** with the following target URL: `http://<ip address>:6000/api/v1/unified?roomId=<roomId>`

*Replace with the actual server IP and matrix room ID.*

{{< hint info >}}
NOTE: multiple `roomId` parameters can be provided to forward the alert to multiple rooms at once
{{< /hint >}}

![screenshot of grafana contact point setup](img/grafanaUnifiedContactPoint.png)

## Step 2

Set up a new **Notification Policy** to forward alerts to the new contact point.

The screenshot below shows it set as the *root policy* - the default for all alerts.

![screenshot of grafana notification policy setup](img/grafanaUnifiedPolicy.png)

## Step 3

Set a **Summary** field to the alert details to have it shown in the Matrix message.

![screenshot of grafana alert details](img/grafanaUnifiedDetails.png)
