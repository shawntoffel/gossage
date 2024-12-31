package gossage

import (
	"database/sql"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

func Db() (*sql.DB, error) {
	dbURL := os.Getenv("GOSSAGE_TEST_CONNECTION_STRING")
	if len(dbURL) < 1 {
		dbURL = "postgresql://root@localhost:26257/defaultdb?sslmode=disable"
	}
	return sql.Open("postgres", dbURL)
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

	err = gossage.RegisterMigrations(migration1{}, migration2{})
	if err != nil {
		t.Error(err)
		return
	}

	err = gossage.Up()
	if err != nil {
		t.Error(err)
		return
	}

	err = gossage.DownTo("0001_migration1")
	if err != nil {
		t.Error(err)
		return
	}
}

type migration1 struct{}

func (m migration1) Version() string {
	return "0001_migration1"
}
func (m migration1) Up(tx *sql.Tx) error {
	_, err := tx.Exec(`CREATE TABLE test1 (id UUID PRIMARY KEY NOT NULL)`)
	return err
}
func (m migration1) Down(tx *sql.Tx) error {
	_, err := tx.Exec(`DROP TABLE test1`)
	return err
}

type migration2 struct{}

func (m migration2) Version() string {
	return "0002_migration2"
}
func (m migration2) Up(tx *sql.Tx) error {
	_, err := tx.Exec(`CREATE TABLE test2 (id UUID PRIMARY KEY NOT NULL)`)
	return err
}
func (m migration2) Down(tx *sql.Tx) error {
	_, err := tx.Exec(`DROP TABLE test2`)
	return err
}
