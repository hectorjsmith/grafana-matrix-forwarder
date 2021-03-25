---
title: Docker
weight: 40
---

The tool can be run inside a docker container if desired.

The docker image can be pulled from the Gitlab registry for this project

```
registry.gitlab.com/hectorjsmith/grafana-matrix-forwarder:latest
```

Check the [registry page](https://gitlab.com/hectorjsmith/grafana-matrix-forwarder/container_registry/1616723) for all available tags.

{{< tabs "dockerExamples" >}}
{{< tab "Docker run" >}}

{{< highlight bash "linenos=table" >}}
docker run -d \
    --name "grafana-matrix-forwarder" \
    -e GMF_MATRIX_USER=@user:matrix.org \
    -e GMF_MATRIX_PASSWORD=password \
    -e GMF_MATRIX_HOMESERVER=matrix.org \
    registry.gitlab.com/hectorjsmith/grafana-matrix-forwarder:latest
{{< /highlight >}}

{{< /tab >}}
{{< tab "Docker Compose" >}}

{{< highlight yaml "linenos=table" >}}
version: "2"
services:
    forwarder:
        image: registry.gitlab.com/hectorjsmith/grafana-matrix-forwarder:latest
        environment:
        - GMF_MATRIX_USER=@user:matrix.org
        - GMF_MATRIX_PASSWORD=password
        - GMF_MATRIX_HOMESERVER=matrix.org
        ports:
- "6000:6000"
{{< /highlight >}}

{{< /tab >}}
{{< /tabs >}}
