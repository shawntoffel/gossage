package gossage

import (
	"database/sql"
	"time"
)

type Version struct {
	Id      string
	Version string
	Created time.Time
}

type MigrationHistory struct {
	db *sql.DB
}

func NewMigrationHistory(db *sql.DB) *MigrationHistory {
	return &MigrationHistory{db: db}
}

func (mh *MigrationHistory) Initialize() error {
	q := `CREATE TABLE IF NOT EXISTS gossage_migration_history (
			id uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
			version STRING UNIQUE NOT NULL,
			created TIMESTAMPTZ NOT NULL DEFAULT NOW())`

	_, err := mh.db.Exec(q)

	return err
}

func (mh *MigrationHistory) AddMigration(m Migration) error {
	q := `INSERT INTO gossage_migration_history (version) VALUES ($1)`

	_, err := mh.db.Exec(q, m.Version())

	return err
}

func (mh *MigrationHistory) LatestVersion() (Version, error) {
	version := Version{}

	q := `SELECT id, version, created FROM gossage_migration_history
			ORDER BY created DESC LIMIT 1`

	err := mh.db.QueryRow(q).Scan(&version.Id, &version.Version, &version.Created)

	return version, err
}
