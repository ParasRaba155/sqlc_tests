package main

import (
	"embed"
	"errors"
	"flag"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

var migrationPath = "migrations"

//go:embed migrations/*.sql
var fs embed.FS

var (
	force   = flag.Bool("force", false, "for setting dirty to non dirty migration forcefully")
	gotoVer = flag.Uint("goto", 0, "for migrating to particular version")
	getVer  = flag.Bool("get", false, "for getting the version of migration")
)

const db_url = `postgres://sqlc-test:pass@localhost:5432/sqlc-test?sslmode=disable`

func main() {
	flag.Parse()

	driver, err := iofs.New(fs, migrationPath)
	if err != nil {
		log.Fatalf("could not create migration driver: %s", err)
	}

	m, err := migrate.NewWithSourceInstance("iofs", driver, db_url)
	if err != nil {
		log.Fatalf("could not create migration source instance: %s", err)
	}

	if *getVer {
		logMigrationVersion(m)
		return
	}

	if *force {
		handleForcedMigrations(m)
		return
	}

	if *gotoVer != 0 {
		logMigrationVersion(m)
		err := m.Migrate(*gotoVer)
		if err != nil {
			log.Fatalf("could not migrate to version %d : %s", *gotoVer, err)
		}
		logMigrationVersion(m)
		return
	}

	logMigrationVersion(m)

	if err = m.Up(); err != nil {
		log.Fatalf("could not migrate up: %s", err)
	}

	log.Printf("successfully migrated the source")
	logMigrationVersion(m)
}

func handleForcedMigrations(m *migrate.Migrate) {
	var version int
	var confirm string

	fmt.Printf("enter the version which you would like to force migration:")
	fmt.Scanf("%d", &version)
	fmt.Printf("are you sure (yes/no):")
	fmt.Scanf("%s", &confirm)
	if confirm[0:1] != "y" {
		return
	}
	logMigrationVersion(m)
	err := m.Force(version)
	if err != nil {
		log.Fatalf("could not force the migration: %s", err)
	}
	logMigrationVersion(m)
}

func logMigrationVersion(m *migrate.Migrate) {
	version, dirty, err := m.Version()
	if err != nil {
		if !errors.Is(err, migrate.ErrNilVersion) {
			log.Fatalf("could not get the current migration version: %s", err)
		}
	}
	log.Printf("migration information :\nversion: %d\nis_dirty: %t", version, dirty)
}
