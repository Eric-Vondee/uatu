package uatu

// Swagger generation
//
// The generated swagger/ package is committed, because server/routes.go
// imports it — the server will not build without it.
//
//go:generate swag init --output swagger -g swagger.go --parseInternal --exclude ./internal/contracts
