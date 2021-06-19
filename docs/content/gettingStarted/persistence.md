---
title: Persistence
weight: 60
---

By default, the tool will create a `grafanaToMatrixMap.json` file on each forwarded alert. This file contains a map from grafana alert to matrix message.

This file is only a copy of the tool's internal state to support restarts. It does not need to be backed up.

{{< hint info >}}
The creation of this file can be disabled by using the `--persistAlertMap` CLI flag or the `GMF_PERSIST_ALERT_MAP` environment variable on startup.
{{< /hint >}}

## Technical Explanation

{{< hint info >}}
The following only applies when the `resolveMode` is set to `reply`.
{{< /hint >}}

To support resolving alerts using a reply, the tool needs to keep track of the original Matrix messages it sent for a given alert. This includes the message ID and the message content.

When a new grafana alert comes in for the same ID and the `resolved` state, the tool tries to find the message it sent for the original alert (with the `alerting` state). If the tool finds a message, it can use that message to generate a reply.

As the map is stored in memory, it is lost when the tool restarts. By persisting the map to a file, the internal state can be restored on startup.

When the tool fails to find an entry for the grafana alert ID, it defaults to send a standard message - even when the resolve mode is `reply`.
