package postgresql

import (
	"database/sql"
	"embed"
	"fmt"
	"log/slog"
	"strings"

	_ "github.com/lib/pq"
)

//go:embed migrations/*
var f embed.FS

func (p *Plugin) initConnection() error {
	// p.createDb()

	// TODO: using plaintext connection to postgres for now, this needs to be configurable
	// connect to the real OpenTalaria database here
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		p.config.Host,
		p.config.Port,
		p.config.Username,
		p.config.Password,
		p.config.Database)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		// try to create the database, if that fails, return the original error and log the create error.
		err1 := p.createDb()
		if err1 != nil {
			slog.Error("error creating pg database", "err", err1)
			return err
		}
	}

	return nil
}

func (p *Plugin) createDb() error {
	// TODO: using plaintext connection to postgres for now, this needs to be configurable
	// connect to the default postgres database first to check if we need to create the OpenTalaria database and tables
	psqlInit := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=postgres sslmode=disable",
		p.config.Host,
		p.config.Port,
		p.config.Username,
		p.config.Password)

	dbInit, err := sql.Open("postgres", psqlInit)
	if err != nil {
		return err
	}
	defer dbInit.Close()

	err = dbInit.Ping()
	if err != nil {
		return err
	}

	createDb, _ := f.ReadFile("migrations/create_db.sql")
	// This is really not recommended, but I have no other way to dynamically create the database off a config file.
	// Here we will rely on the fact that the database name is not a user input, but rather the admin who hosts OpenTalaria has to specify in in the config.
	// Admins please don't sql inject malicious statements via the config file.
	prepared := strings.Replace(string(createDb), "{dbname}", p.config.Database, 1)
	dbInit.Exec(prepared)

	return nil
}
