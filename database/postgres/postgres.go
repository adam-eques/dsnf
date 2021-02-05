package postgres

import (
	"fmt"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

type Config struct {
	Host     string
	Password string
	User     string
	Port     string
	DBName   string
	SSLMode  string
	DBUrl    string
}

func SetupDatabase(db *gorm.DB, models ...interface{}) error {
	err := db.AutoMigrate(models...)
	return err
}

func New(config Config) (*gorm.DB, error) {
	var (
		db  *gorm.DB
		err error
	)

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, err
}
