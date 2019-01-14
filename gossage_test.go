package gossage

import (
	"database/sql"
	"testing"

	_ "github.com/lib/pq"
)

func Db() (*sql.DB, error) {
	return sql.Open("postgres", "postgresql://test@localhost:26257/test?sslmode=disable")
}

func TestMigrationHistory(t *testing.T) {
	db, err := Db()
	if err != nil {
		t.Error(err)
		return
	}

	gossage, err := New(db)
	if err != nil {
		t.Error(err)
		return
	}

	gossage.RegisterMigration(migration1{})

	err = gossage.Up()
	if err != nil {
		t.Error(err)
		return
	}
}

type migration1 struct{}

func (m migration1) Version() string {
	return "0001"
}
func (m migration1) Up(tx *sql.Tx) error {
	return nil
}
func (m migration1) Down(tx *sql.Tx) error {
	return nil
}
