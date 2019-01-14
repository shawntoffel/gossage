package gossage

import (
	"database/sql"
)

type Migration interface {
	Version() string
	Up(*sql.Tx) error
	Down(*sql.Tx) error
}

type ByVersion []Migration

func (v ByVersion) Len() int {
	return len(v)
}

func (v ByVersion) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}

func (v ByVersion) Less(i, j int) bool {
	return v[i].Version() < v[j].Version()
}
