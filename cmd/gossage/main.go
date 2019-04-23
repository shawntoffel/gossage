package main

import (
	"fmt"
	"os"

	"github.com/shawntoffel/gossage/internal/cli"
)

func main() {
	migrationName := ""

	if len(os.Args) > 1 {
		migrationName = os.Args[1]
	}

	err := cli.GenerateMigrationFile(migrationName)
	if err != nil {
		fmt.Printf("gossage: could not generate migration file: %s", err)
	}
}
