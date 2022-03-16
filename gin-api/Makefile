.PHONY: setup doc lint vet\
		network start stop down\
		build migrate\
		test

setup:
	sudo pip install pre-commit
	go install github.com/kyoh86/richgo@latest
	go install -v golang.org/x/tools/cmd/godoc@latest
	./scripts/install-linters.sh
	pre-commit install

# Utility commands
doc:
	godoc -http=:6061

lint:
	gofmt -s -w .
	goimports -w .
	-errcheck ./...
	-go vet --vettool=$$(which shadow) ./...

vet:
	go vet --vettool=$$(which shadow) ./...

# Test commands
test:
	go clean -testcache
	-richgo test -tags=unit,integration -p 1 ./apps/...

# Development commands
start:
	go run main.go http