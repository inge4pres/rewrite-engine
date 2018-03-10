test:
	go vet ./...
	go test -v ./pkg/... ./cmd/...

build: test
	GOOS=linux GOARCH=amd64 go build -o rewrite-engine ./cmd/rewrite-engine

package: build
	docker build -t registry.bravofly.intra:5000/application/rewrite-engine .

run-docker:
	docker run --rm -p 8081:8081 registry.bravofly.intra:5000/application/rewrite-engine 

.PHONY: test build package run-docker