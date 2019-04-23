package cli

import (
	"html/template"
	"os"
	"strings"
	"time"
)

type migration struct {
	Name      string
	Timestamp string
}

// Generates a migration go file via template. Expects a snake_case name.
func GenerateMigrationFile(name string) error {
	t, err := template.New("generate").Parse(migrationTemplate)
	if err != nil {
		return err
	}

	m := migration{
		Name:      toCamelCase(name),
		Timestamp: time.Now().Format("20060102150405"),
	}

	file, err := os.Create(m.Timestamp + "_" + name + ".go")
	if err != nil {
		return err
	}

	defer file.Close()

	err = t.Execute(file, m)
	if err != nil {
		return err
	}

	return nil
}

func toCamelCase(name string) string {
	camelCase := ""
	capitalizeNextChar := true

	for _, char := range name {
		c := string(char)

		if capitalizeNextChar {
			c = strings.ToUpper(c)
		}

		if c == "_" {
			capitalizeNextChar = true
		} else {
			capitalizeNextChar = false
			camelCase += c
		}
	}

	return camelCase
}

const migrationTemplate = `package migrations

import "database/sql"

type {{.Name}}{{.Timestamp}} struct{}

func (m {{.Name}}{{.Timestamp}}) Version() string {
	return "{{.Timestamp}}_{{.Name}}"
}

func (m {{.Name}}{{.Timestamp}}) Up(tx *sql.Tx) error {
	_, err := tx.Exec(` + "``" + `)
	return err
}

func (m {{.Name}}{{.Timestamp}}) Down(tx *sql.Tx) error {
	_, err := tx.Exec(` + "``" + `)
	return err
}`
