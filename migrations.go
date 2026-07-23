package uatu

import "embed"

// MigrationsDir is the path to Migrations inside the embedded filesystem.
const MigrationsDir = "internal/storage/postgres/migrations"

//go:embed internal/storage/postgres/migrations/*.sql
var Migrations embed.FS
