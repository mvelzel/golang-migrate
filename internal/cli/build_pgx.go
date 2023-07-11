//go:build pgx
// +build pgx

package cli

import (
	_ "github.com/mvelzel/golang-migrate/v4/database/pgx"
)
