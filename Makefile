install-deps:
	go mod download

# Standard go test
test:
	go test ./... -v -race

# Make sure no unnecessary dependencies are present
go-mod-tidy:
	go mod tidy -v
	git diff-index --quiet HEAD

format:
	go fmt $(go list ./... | grep -v /vendor/)
	go vet $(go list ./... | grep -v /vendor/)

define prepare_build_vars
    $(eval VERSION_FLAG=-X 'main.appVersion=$(shell git describe --tags)')
endef

build/local:
	$(call prepare_build_vars)
	go build -a --ldflags "${VERSION_FLAG}" -o build/grafana-matrix-forwarder.bin ./app.go
