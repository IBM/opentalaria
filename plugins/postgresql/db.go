package postgresql

import (
	"database/sql"
	"embed"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
)

const DB_NAME = "opentalaria"

func (p *Plugin) getConnectionString(db string) string {
	return fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable binary_parameters=yes",
		p.config.Host,
		p.config.Port,
		p.config.Username,
		p.config.Password,
		db)

}

func (p *Plugin) initConnection() error {
	// TODO: using plaintext connection to postgres for now, this needs to be configurable
	// connect to the real OpenTalaria database here

	db, err := sql.Open("postgres", p.getConnectionString(DB_NAME))
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		// try to create the database, if that fails, return the original error and the create error.
		err1 := p.createDb()
		if err1 != nil {
			return errors.Join(err, err1)
		}
	}

	// TODO: implement a migration process.

	p.db = db

	return nil
}

//go:embed migrations/*
var f embed.FS

func (p *Plugin) createDb() error {
	// TODO: using plaintext connection to postgres for now, this needs to be configurable
	// connect to the default postgres database first to check if we need to create the OpenTalaria database and tables
	dbInit, err := sql.Open("postgres", p.getConnectionString("postgres"))
	if err != nil {
		return err
	}
	defer dbInit.Close()

	err = dbInit.Ping()
	if err != nil {
		return err
	}

	_, err = dbInit.Exec("create database " + DB_NAME)
	if err != nil {
		return err
	}

	// TODO: using plaintext connection to postgres for now, this needs to be configurable
	// connect to the opentalaria postgres database and run the migration scripts
	dbMigration, err := sql.Open("postgres", p.getConnectionString(DB_NAME))
	if err != nil {
		return err
	}
	defer dbMigration.Close()

	initDbScript, _ := f.ReadFile("migrations/init.sql")
	_, err = dbMigration.Exec(string(initDbScript))
	if err != nil {
		return err
	}

	return nil
}
