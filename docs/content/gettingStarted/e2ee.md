---
title: End to End Encryption
weight: 50
---

This tool does not natively support sending alerts to matrix rooms with encryption enabled.

However, encrypted rooms are supported by using [pantalaimon](https://github.com/matrix-org/pantalaimon) to act as a reverse proxy that handles the encryption.
Information on setting up pantalaimon can be found on the project's Github page.

This tool can be configured to forward messages through the pantalaimon proxy server by setting the `GMF_MATRIX_HOMESERVER` environment variable or `-homeserver` cli argument to your pantalaimon instance.

For example:
`-homeserver http://localhost:6000`
