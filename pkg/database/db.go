package database

import (
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate"
	pg "github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file" // this is needed to migrate from a file

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB(cfg *DatabaseConfig) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Berlin", cfg.Host, cfg.User, cfg.Password, cfg.Name, cfg.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("could not connect to db")
	}

	return db
}

func MigrateUp(db *gorm.DB, cfg *DatabaseConfig) {
	sqlDB, err := db.DB()

	if err != nil {
		panic(err)
	}

	driver, err := pg.WithInstance(sqlDB, &pg.Config{
		MigrationsTable: "schema_migrations",
	})

	if err != nil {
		panic(err)
	}

	mig, err := migrate.NewWithDatabaseInstance("file://"+cfg.MigrationDir, cfg.Name, driver)
	if err != nil {
		panic(err)
	}

	if err = mig.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		panic(err)
	}
}
