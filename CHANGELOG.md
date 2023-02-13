# CHANGELOG

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog], and this project adheres to [Semantic Versioning].

## [Unreleased]

## [0.7.0] - 2023-02-13
### Features
- (b267c16) feat: add new raw data field to message
- (aad2690) feat: support multiple alerts at once
- (56e1f06) feat: support unified alerts ([#33](https://gitlab.com/hectorjsmith/grafana-matrix-forwarder/issues/33))

### Bug Fixes
- (7f95815) fix: html format of raw data
- (3e2c76a) fix: formatting of matrix replies
- (30d2894) fix: error in metric collector

### Merge Requests
- (82c6e8e) Merge branch 'chore/dependency-updates' into 'main'
- (d7a0224) Merge branch 'chore/update-matrix-client-version' into 'main'
- (bac27a1) Merge branch 'fix-html-format-of-raw-data' into 'main'
- (d7b60de) Merge branch 'add-new-raw-data-field-to-message' into 'main'
- (22555e0) Merge branch 'fix-formatting-of-matrix-replies' into 'main'
- (be1c44f) Merge branch 'support-multiple-alerts-in-a-single-webhook' into 'main'
- (cfe814f) Merge branch 'rename-internal-alert-data-struct' into 'main'
- (86febe8) Merge branch '33-support-new-json-format-from-grafana' into 'main'
- (4dd3c39) Merge branch 'split-formatter-and-forwarder-code' into 'main'
- (55626c6) Merge branch 'fix-error-in-metrics-collector' into 'main'
- (0fc40a6) Merge branch 'split-request-handler-code-into-shared-and-version-specific-packages' into 'main'
- (b221749) Merge branch 'update-metric-collector-to-handle-any-alert-state' into 'main'
- (3e957c8) Merge branch 'add-model-package-for-shared-structs' into 'main'
- (e439b64) Merge branch 'better-handling-of-forwarding-alert-to-multiple-rooms' into 'main'
- (50a6b9b) Merge branch 'only-publish-stable-images-to-the-latest-tag' into 'main'
- (0926489) Merge branch 'move-metric-collection-code-to-new-package' into 'main'
- (d8a185a) Merge branch 'refactor-forwarder-to-use-data-struct-instead-of-raw-payload' into 'main'
- (a758137) Merge branch 'move-payload-structs-to-new-v0-package' into 'main'
- (269d742) Merge branch '31-add-prometheus-library-for-metrics' into 'main'
- (5cceebd) Merge branch 'release/0.6.0' into 'main'


## [0.6.0] - 2021-06-22
### Features
- (c9adf61) feat: option to persist alert map ([#30](https://gitlab.com/hectorjsmith/grafana-matrix-forwarder/issues/30))
- (7e592e6) feat: support saving event map to file ([#24](https://gitlab.com/hectorjsmith/grafana-matrix-forwarder/issues/24))
- (4499c99) feat: support forwarding alert to multiple rooms ([#10](https://gitlab.com/hectorjsmith/grafana-matrix-forwarder/issues/10))
- (26f66ef) feat: include global tags in forwarded alert ([#26](https://gitlab.com/hectorjsmith/grafana-matrix-forwarder/issues/26))

### Merge Requests
- (8136ee1) Merge branch '30-option-to-disable-persisting-alert-to-matrix-message-map' into 'main'
- (8cfdc2d) Merge branch 'refactor/split-forwarder-code' into 'main'
- (4e22239) Merge branch 'refactor/split-forwarder-into-new-folder' into 'main'
- (4e7e960) Merge branch '24-persistent-alert-to-message-event-map' into 'main'
- (684f192) Merge branch '10-support-multiple-room-ids-in-the-url' into 'main'
- (4d70b48) Merge branch '14-support-for-e2ee-rooms' into 'main'
- (9e1601d) Merge branch '26-include-tags-in-forwarded-alert' into 'main'
- (357ca58) Merge branch 'release/v0.5.0' into 'main'


## [0.5.0] - 2021-03-27
### Features
- (6b552e0) feat: cli flags overwrite environment variables
- (c4e53d5) feat: support for rounding metric values ([#29](https://gitlab.com/hectorjsmith/grafana-matrix-forwarder/issues/29))
- (b880a0c) feat: new cli flag to set value rounding ([#29](https://gitlab.com/hectorjsmith/grafana-matrix-forwarder/issues/29))
- (37572f3) feat: new -env cli flag to use environment variables ([#27](https://gitlab.com/hectorjsmith/grafana-matrix-forwarder/issues/27))
- (01736eb) feat: include metric values in alerts ([#8](https://gitlab.com/hectorjsmith/grafana-matrix-forwarder/issues/8))
- (8368043) feat: add missing fields to go struct ([#8](https://gitlab.com/hectorjsmith/grafana-matrix-forwarder/issues/8))

### Merge Requests
- (b2e7491) Merge branch 'docs/add-missing-cli-and-env-flag' into 'main'
- (6d3f30c) Merge branch 'docs/fix-main-branch-in-badge-links' into 'main'
- (8164368) Merge branch 'docs/structure-documentation-site' into 'main'
- (79c6a8a) Merge branch 'refactor/update-grafana-forwarder-handler' into 'main'
- (42da71a) Merge branch 'refactor/improve-cli-parsing' into 'main'
- (89e1283) Merge branch '29-round-metric-values' into 'main'
- (865fec5) Merge branch 'ci/fix-broken-build-pipeline' into 'main'
- (87c6291) Merge branch 'build/tidy-makefile-and-gitlab-ci' into 'main'
- (c696e76) Merge branch '28-create-a-documentation-site' into 'main'
- (9c7e68b) Merge branch '27-improve-support-for-environment-variables' into 'main'
- (c4a5460) Merge branch 'refactor/use-single-template-for-alert-messages' into 'main'
- (2bd5b5f) Merge branch '8-include-metric-values-in-alert' into 'main'
- (90cf633) Merge branch '23-build-release-version-for-tag' into 'main'
- (9a6c2b1) Merge branch '25-move-all-code-to-new-source-folder' into 'main'
- (32c5ff6) Merge branch 'release/v0.4.0' into 'main'


## [0.4.0] - 2021-01-10
### Bug Fixes
- (517c066) fix: do not show log for resolve mode in version mode ([#21](https://gitlab.com/hectorjsmith/grafana-matrix-forwarder/issues/21))

### Features
- (761aedd) feat: support resolving alerts with replies ([#20](https://gitlab.com/hectorjsmith/grafana-matrix-forwarder/issues/20))
- (6281bf3) feat: add dockerfile to run forwarder using docker ([#5](https://gitlab.com/hectorjsmith/grafana-matrix-forwarder/issues/5))

### Merge Requests
- (5f8513c) Merge branch 'fix/do-not-escape-html-in-reply-body' into 'main'
- (d43d3e7) Merge branch '20-resolve-alerts-with-reply' into 'main'
- (0d4ccd4) Merge branch '5-add-docker-image' into 'main'
- (e4fbb82) Merge branch 'docs/fix-typo-in-readme' into 'main'
- (01ee605) Merge branch '21-resolve-mode-showing-when-printing-version' into 'main'
- (6af9079) Merge branch 'release/v0.3.0' into 'main'


## [0.3.0] - 2020-12-14
### Features
- (9ef976e) feat: export metrics on alert state ([#17](https://gitlab.com/hectorjsmith/grafana-matrix-forwarder/issues/17))
- (b86caa0) feat: export forward count metrics ([#17](https://gitlab.com/hectorjsmith/grafana-matrix-forwarder/issues/17))
- (078cace) feat: add metrics endpoint ([#17](https://gitlab.com/hectorjsmith/grafana-matrix-forwarder/issues/17))
- (0399f02) feat: support resolving alerts with reactions ([#11](https://gitlab.com/hectorjsmith/grafana-matrix-forwarder/issues/11))
- (5cf2fdb) feat: load id fields from alert payload ([#11](https://gitlab.com/hectorjsmith/grafana-matrix-forwarder/issues/11))
- (aa30c24) feat: support for sending reactions ([#11](https://gitlab.com/hectorjsmith/grafana-matrix-forwarder/issues/11))

### Merge Requests
- (2b8354d) Merge branch '17-expose-prometheus-metrics-for-forward-count' into 'main'
- (15473ad) Merge branch '18-only-run-full-build-on-tags' into 'main'
- (9d74d7a) Merge branch '11-resolve-alerts-with-a-reaction-reply' into 'main'
- (477c447) Merge branch '16-use-a-go-template-for-the-matrix-message-format' into 'main'
- (523ef99) Merge branch 'release/v0.2.0' into 'main'


## [0.2.0] - 2020-12-04
### Bug Fixes
- (feee9ea) fix: add messenger file
- (dddf6e7) fix: strip html tags from matrix message body

### Features
- (6555459) feat: support no_data alert states
- (c55cb8b) feat: support no_data alert states
- (e507f08) feat: new cli option to log alert payloads

### Merge Requests
- (b1b7197) Merge branch '15-add-go-report-card' into 'main'
- (b0f1731) Merge branch 'agile/rewrite-readme-file' into 'main'
- (d796b8b) Merge branch '12-create-compressed-file-with-build-artefacts' into 'main'
- (e3d09d1) Merge branch 'release/v0.1.0' into 'main'


## [0.1.0] - 2020-11-22
### Bug Fixes
- (252ac7d) fix: handle unknown alert states
- (ce963d0) fix: handle alert resolution messages

### Features
- (2422339) feat: add cli flags for host and port
- (58a88a6) feat: include rule name in alert

### Merge Requests
- (13405f6) Merge branch '2-write-basic-readme-file-with-instructions' into 'main'
- (0547b56) Merge branch '7-correctly-handle-resolved-alerts' into 'main'
- (a0df35e) Merge branch '3-configure-server-port-and-address-on-startup' into 'main'
- (186832e) Merge branch '1-setup-basic-gitlab-ci-pipeline' into 'main'
- (e120688) Merge branch '4-include-rule-name-in-alert' into 'main'

---

*This changelog is automatically generated by [git-chglog]*

[Keep a Changelog]: https://keepachangelog.com/en/1.0.0/
[Semantic Versioning]: https://semver.org/spec/v2.0.0.html
[git-chglog]: https://github.com/git-chglog/git-chglog
[Unreleased]: https://gitlab.com/hectorjsmith/grafana-matrix-forwarder/compare/0.7.0...main
[0.7.0]: https://gitlab.com/hectorjsmith/grafana-matrix-forwarder/compare/0.6.0...0.7.0
[0.6.0]: https://gitlab.com/hectorjsmith/grafana-matrix-forwarder/compare/0.5.0...0.6.0
[0.5.0]: https://gitlab.com/hectorjsmith/grafana-matrix-forwarder/compare/0.4.0...0.5.0
[0.4.0]: https://gitlab.com/hectorjsmith/grafana-matrix-forwarder/compare/0.3.0...0.4.0
[0.3.0]: https://gitlab.com/hectorjsmith/grafana-matrix-forwarder/compare/0.2.0...0.3.0
[0.2.0]: https://gitlab.com/hectorjsmith/grafana-matrix-forwarder/compare/0.1.0...0.2.0
[0.1.0]: https://gitlab.com/hectorjsmith/grafana-matrix-forwarder/compare/0.0.0...0.1.0
