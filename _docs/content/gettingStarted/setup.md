---
title: Run Binary
weight: 20
---

The tool is compiled as a single binary, so it is simple to run as a standalone application.

To run the forwarder just provide credentials to connect to a matrix account to send messages from.

```
$ ./grafana-matrix-forwarder --user @userId:matrix.org --password xxx --homeserver matrix.org
```

See the page on [configuration]({{< ref "configuration.md" >}}) to change the port the forwarder listens on.
