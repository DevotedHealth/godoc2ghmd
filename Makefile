.PHONY: all lint test

all:
	go generate

test:
	go test -cover .

lint:
	@echo "Checking if there are files that need gofmt (goimports) to be run"
	@ ! find . -type f -name '*.go' -not -path "*/vendor/*" -not -path "*/proto/*" -not -path "*generated*.go" -print0 | xargs -0 goimports -l | sed 's/^/Needs reformatting: /'| grep .
	@echo "Running gometalinter"
	@gometalinter --config=.gometalinter_config.json ./...
