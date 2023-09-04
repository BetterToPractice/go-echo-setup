package lib

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Database struct {
	ORM *gorm.DB
}

func NewDatabase(config Config, logger Logger) Database {
	dbConfig := postgres.Config{
		DSN: config.Database.DSN(),
	}

	db, err := gorm.Open(postgres.New(dbConfig), &gorm.Config{
		SkipDefaultTransaction:                   true,
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		QueryFields: true,
	})

	if err != nil {
		logger.Zap.Fatalf("Error to open database [%s] connection: %v", dbConfig.DSN, err)
	}

	logger.Zap.Info("Database connection established")
	return Database{
		ORM: db,
	}
}
