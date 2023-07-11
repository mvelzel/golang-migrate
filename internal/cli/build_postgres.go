//go:build postgres
// +build postgres

package cli

import (
	_ "github.com/mvelzel/golang-migrate/v4/database/postgres"
)
