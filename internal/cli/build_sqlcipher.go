//go:build sqlcipher
// +build sqlcipher

package cli

import (
	_ "github.com/mvelzel/golang-migrate/v4/database/sqlcipher"
)
