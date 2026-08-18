.PHONY: build run seed migrate docs docs-check tidy vet


build:
	go build -o bin/uatu     ./cmd
	go build -o bin/uatu-cli ./cmd/cli

run:
	go run ./cmd

migrate:
	go run ./cmd/cli migrate

seed:
	go run ./cmd/cli seed

# Regenerate the OpenAPI spec from the handler annotations. The generated
# swagger/ package is committed, because server/routes.go imports it — the
# server will not build without it.
docs:
	go generate ./...

# Fails if the committed spec has drifted from the annotations. Suitable for CI.
docs-check:
	@go generate ./...
	@git diff --exit-code -- swagger/ \
		|| (echo "\nswagger/ is stale — run 'make docs' and commit the result." && exit 1)

tidy:
	go mod tidy

vet:
	go vet ./...
