package gossage

import "fmt"

type Gossage struct {
	migrations map[string]Migration
}

func (g Gossage) RegisterMigration(m Migration) error {
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

func (g Gossage) Up() error {

	return nil
}

func (g Gossage) Down() error {

	return nil
}
