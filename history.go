package gossage

import (
	"database/sql"
	"time"
)

type Version struct {
	Id      string
	Version string
	Applied bool
	Created time.Time
	Updated time.Time
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
			applied BOOL NOT NULL DEFAULT false,
			created TIMESTAMPTZ NOT NULL DEFAULT NOW(),
			updated TIMESTAMPTZ
		)`

	_, err := mh.db.Exec(q)

	return err
}

func (mh *MigrationHistory) AddMigration(m Migration) error {
	q := `INSERT INTO gossage_migration_history (version, applied) VALUES ($1, true)
			ON CONFLICT (version) DO UPDATE SET applied = true, updated = NOW()`

	_, err := mh.db.Exec(q, m.Version())

	return err
}

func (mh *MigrationHistory) LatestVersion() (Version, error) {
	version := Version{}

	q := `SELECT id, version, created FROM gossage_migration_history 
			WHERE applied = true
			ORDER BY version DESC LIMIT 1`

	err := mh.db.QueryRow(q).Scan(&version.Id, &version.Version, &version.Created)

	return version, err
}

func (mh *MigrationHistory) VersionsGreaterThan(version string) ([]string, error) {
	versions := []string{}

	q := `SELECT version FROM gossage_migration_history 
			WHERE applied = true
			AND version > $1
			ORDER BY version`

	rows, err := mh.db.Query(q, version)
	if err != nil {
		return versions, err
	}

	defer rows.Close()
	for rows.Next() {
		v := ""
		err := rows.Scan(&v)
		if err != nil {
			return versions, err
		}

		versions = append(versions, v)
	}

	return versions, nil
}

func (mh *MigrationHistory) RevertMigration(version string) error {
	q := `UPDATE gossage_migration_history
			SET applied = false
			WHERE version = $1`

	_, err := mh.db.Exec(q, version)

	return err
}
