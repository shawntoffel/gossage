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

	err = gossage.RegisterMigration(migration1{})
	if err != nil {
		t.Error(err)
		return
	}
	err = gossage.RegisterMigration(migration2{})
	if err != nil {
		t.Error(err)
		return
	}

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

type migration2 struct{}

func (m migration2) Version() string {
	return "0002"
}
func (m migration2) Up(tx *sql.Tx) error {
	return nil
}
func (m migration2) Down(tx *sql.Tx) error {
	return nil
}
