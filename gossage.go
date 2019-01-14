package gossage

import (
	"database/sql"
	"fmt"
	"sort"
)

type Gossage struct {
	db         *sql.DB
	migrations map[string]Migration
	history    *MigrationHistory
}

func New(db *sql.DB) (*Gossage, error) {
	g := &Gossage{
		db:         db,
		migrations: map[string]Migration{},
		history:    NewMigrationHistory(db),
	}

	err := g.history.Initialize()
	return g, err
}

func (g *Gossage) RegisterMigration(m Migration) error {
	if g.migrations == nil {
		g.migrations = map[string]Migration{}
	}

	version := m.Version()

	_, exists := g.migrations[version]
	if exists {
		return fmt.Errorf("gossage: a migration with version '%s' has already been registered.", version)
	}

	g.migrations[version] = m

	return nil
}

func (g *Gossage) Up() error {
	current, err := g.history.LatestVersion()
	if err != nil && err != sql.ErrNoRows {
		return fmt.Errorf("gossage: could not determine current migration version. %s", err)
	}

	migrations := []Migration{}

	for _, val := range g.migrations {
		if val.Version() > current.Version {
			migrations = append(migrations, val)
		}
	}

	if len(migrations) < 1 {
		msglog("no migrations to perform. Version: %s", current.Version)
		return nil
	}

	sort.Sort(ByVersion(migrations))

	for _, m := range migrations {
		tx, err := g.db.Begin()
		if err != nil {
			return err
		}

		err = m.Up(tx)
		if err != nil {
			return err
		}

		err = tx.Commit()
		if err != nil {
			return err
		}

		g.history.AddMigration(m)
		msglog("completed migration '%s'", m.Version())
	}

	msglog("migrations complete")

	return nil
}

func (g *Gossage) Down() error {

	return nil
}
