all: vet fmt build

build:
	docker build -f build-tools/Dockerfile -t f5ipmanager .

vet:
	go vet ./...

fmt:
	go list -f '{{.Dir}}' ./... | grep -v /vendor/ | xargs -L1 gofmt -l

tidy:
	go mod tidy

vendor:
	go mod vendor
