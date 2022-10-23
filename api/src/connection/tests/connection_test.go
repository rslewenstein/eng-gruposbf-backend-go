package connection_test

import (
	"database/sql"
	"testing"

	"go.mod/src/connection"
)

func TestConnect(t *testing.T) {
	var retDB *sql.DB
	retDB, err := connection.Connect()

	if retDB == nil {
		t.Failed()
	}

	if err != nil {
		t.Failed()
	}
}
