# List make commands
.PHONY: ls
ls:
	cat Makefile | grep "^[a-zA-Z#].*" | cut -d ":" -f 1 | sed s';#;\n#;'g

# Download dependencies
.PHONY: download
download:
	go mod download

# Update project dependencies
.PHONY: update
update:
	go get -u
	go mod download
	go mod tidy

# Run project tests
.PHONY: test
test: download
	go test ./... -v -race

# Look for "suspicious constructs" in source code
.PHONY: vet
vet: download
	go vet ./...

# Format code
.PHONY: fmt
fmt: download
	go mod tidy
	go fmt ./...

# Check for unformatted go code
.PHONY: check/fmt
check/fmt: download
	test -z $(shell gofmt -l .)

# Build project
.PHONY: build
build:
	CGO_ENABLED=0 go build \
	-ldflags "\
	-X main.version=${shell git describe --tags} \
	-X main.commit=${shell git rev-parse HEAD} \
	-X main.date=${shell date --iso-8601=seconds} \
	-X main.builtBy=manual \
	" \
	-o grafana-matrix-forwarder \
	app.go

# Build project docker container
.PHONY: build/docker
build/docker: build
	docker build -t grafana-matrix-forwarder .

# Download theme used for docs site
.PHONY: docs/downloadTheme
docs/downloadTheme:
	wget -O geekdoc.tar.gz https://github.com/thegeeklab/hugo-geekdoc/releases/download/v0.10.1/hugo-geekdoc.tar.gz
	mkdir -p _docs/themes/hugo-geekdoc
	tar -xf geekdoc.tar.gz -C _docs/themes/hugo-geekdoc/

# Build the docs site
.PHONY: docs/build
docs/build:
	cd _docs/ && hugo
