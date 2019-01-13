package gossage

import "database/sql"

type Migration interface {
	Version() string
	Up(*sql.Tx) error
	Down(*sql.Tx) error
}
