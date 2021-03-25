---
title: Documentation
---

<!-- markdownlint-capture -->
<!-- markdownlint-disable MD033 -->

<span class="badge-placeholder">[![Project](https://img.shields.io/badge/project-gitlab-brightgreen?style=flat&logo=gitlab)](https://gitlab.com/hectorjsmith/grafana-matrix-forwarder/)</span>
<span class="badge-placeholder">[![Build Status](https://gitlab.com/hectorjsmith/grafana-matrix-forwarder/badges/main/pipeline.svg)](https://gitlab.com/hectorjsmith/grafana-matrix-forwarder/commits/main)</span>
<span class="badge-placeholder">[![Go Report Card](https://goreportcard.com/badge/gitlab.com/hectorjsmith/grafana-matrix-forwarder)](https://goreportcard.com/report/gitlab.com/hectorjsmith/grafana-matrix-forwarder)</span>
<span class="badge-placeholder">[![License: MIT](https://img.shields.io/badge/license-MIT-brightgreen)](https://gitlab.com/hectorjsmith/csharp-performance-recorder/-/blob/main/LICENSE)</span>

<!-- markdownlint-restore -->
# Grafana to Matrix Forwarder
*Forward alerts from [Grafana](https://grafana.com) to a [Matrix](https://matrix.org) chat room.*

![screenshot of matrix alert message](img/alertExample.png)

* 📦 **Portable**
    * As a single binary the tool is easy to run in any environment
* 📎 **Simple**
    * No config files, all required parameters provided on startup
* 🪁 **Flexible**
    * Support multiple grafana alert channels to multiple matrix rooms
* 📈 **Monitorable**
    * Export metrics to track successful and failed forwards

---

*Made possible by the [maunium.net/go/mautrix](https://maunium.net/go/mautrix/) library and all the contributors to the [matrix.org](https://matrix.org) protocol.*
