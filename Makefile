go/downloadDependencies:
	cd src/ && go mod download

# Standard go test
go/test:
	cd src/ && go test ./... -v -race

# Make sure no unnecessary dependencies are present
go/tidyDependencies:
	cd src/ && go mod tidy -v
	git --no-pager diff
	git diff-index --quiet HEAD

go/format:
	cd src/ && go fmt $(go list ./... | grep -v /vendor/)
	cd src/ && go vet $(go list ./... | grep -v /vendor/)

docs/downloadTheme:
	wget -O geekdoc.tar.gz https://github.com/thegeeklab/hugo-geekdoc/releases/download/v0.10.1/hugo-geekdoc.tar.gz
	mkdir -p docs/themes/hugo-geekdoc
	tar -xf geekdoc.tar.gz -C docs/themes/hugo-geekdoc/

docs/build:
	cd docs/ && hugo

docs/generateChangelog:
	./tools/git-chglog_linux_amd64 --config tools/chglog/config.yml 0.1.0.. > CHANGELOG.md

build/snapshot:
	./tools/goreleaser_linux_amd64 --snapshot --rm-dist --skip-publish

build/release:
	git --no-pager diff
	./tools/goreleaser_linux_amd64 --rm-dist --skip-publish

build/docker:
	docker build -t registry.gitlab.com/hectorjsmith/grafana-matrix-forwarder:latest .

build/dockerTag:
	docker build -t registry.gitlab.com/hectorjsmith/grafana-matrix-forwarder:$(shell git describe --tags) .
