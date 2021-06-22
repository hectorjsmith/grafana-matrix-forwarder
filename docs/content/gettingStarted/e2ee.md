---
title: End to End Encryption
weight: 50
---

This tool does not natively support sending alerts to matrix rooms with encryption enabled.

However, encrypted rooms are supported by using [pantalaimon](https://github.com/matrix-org/pantalaimon) to act as a reverse proxy that handles the encryption.
Information on setting up pantalaimon can be found on the project's Github page.

The grafana alert forwarder can be configured to send messages through the pantalaimon proxy server by setting the `GMF_MATRIX_HOMESERVER` environment variable (or `-homeserver` cli argument) to point at your pantalaimon instance.

## Example

The following docker compose file demonstrates how to run both the forwarder and pantalaimon together.

{{< highlight yaml "linenos=table" >}}
version: "2"
services:
  pantalaimon:
    image: matrixdotorg/pantalaimon
      restart: unless-stopped
      volumes:
      - /docker/pantalaimon:/data

  forwarder:
    image: registry.gitlab.com/hectorjsmith/grafana-matrix-forwarder:latest
    restart: unless-stopped
    environment:
    - GMF_MATRIX_USER=@user:matrix.org
    - GMF_MATRIX_PASSWORD=pw
    - GMF_MATRIX_HOMESERVER=http://pantalaimon:8080
    ports:
    - 6000:6000
{{< /highlight >}}

{{< hint info >}}
NOTE: This assumes that pantalaimon has been configured to run on port `8080`
{{< /hint >}}
