package lib

import (
	"database/sql"
	"errors"
	"github.com/pressly/goose/v3"
	"regexp"
	"strconv"
)

type Migration struct {
	config Config
}

func NewMigration(config Config) Migration {
	return Migration{
		config: config,
	}
}

func (l Migration) Create(filename string) error {
	if err := goose.Run("create", nil, l.config.Database.MigrationDir, filename, "sql"); err != nil {
		return err
	}
	return nil
}

func (l Migration) Migrate(command string, filename string, database Database) error {
	db, err := database.ORM.DB()
	if err != nil {
		return err
	}

	switch command {
	case "up":
		if err := l.Up(filename, db); err != nil {
			return err
		}
	case "down":
		if err := l.Down(filename, db); err != nil {
			return err
		}
	case "redo":
		if err := l.Redo(db); err != nil {
			return err
		}
	}
	return nil
}

func (l Migration) Redo(db *sql.DB) error {
	if err := goose.Redo(db, l.config.Database.MigrationDir); err != nil {
		return err
	}
	return nil
}

func (l Migration) Up(filename string, db *sql.DB) error {
	if filename != "" {
		version, err := GetVersion(filename)
		if err != nil {
			return err
		}
		if err := goose.UpTo(db, l.config.Database.MigrationDir, version); err != nil {
			return err
		}
	} else {
		err := goose.Up(db, l.config.Database.MigrationDir)
		return err
	}

	return nil
}

func (l Migration) Down(filename string, db *sql.DB) error {
	if filename != "" {
		version, err := GetVersion(filename)
		if err != nil {
			return err
		}
		if err := goose.DownTo(db, l.config.Database.MigrationDir, version); err != nil {
			return err
		}
	} else {
		err := goose.Down(db, l.config.Database.MigrationDir)
		return err
	}
	return nil
}

func GetVersion(filename string) (int64, error) {
	re := regexp.MustCompile(`(\d{14})`)
	match := re.FindString(filename)
	if match != "" {
		version, err := strconv.Atoi(match)
		if err == nil {
			return int64(version), err
		}
	}
	return int64(0), errors.New("invalid")
}
