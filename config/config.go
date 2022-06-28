package config

import (
	"database/sql"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type dBConfig struct {
	dbHost     string
	dbPort     string
	dbUser     string
	dbPassword string
	dbName     string
}

type Config struct {
	Db *gorm.DB
}

func (c *Config) initDb() {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	env := "dev"

	dBConfig := dBConfig{
		dbHost, dbPort, dbUser, dbPassword, dbName,
	}
	// urutan url koneksi ke db postgres
	// postgres://dbUser:dbPassword@dbHost:dbPort/dbName
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dBConfig.dbHost, dBConfig.dbUser, dBConfig.dbPassword, dBConfig.dbName, dBConfig.dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	if env == "dev" {
		// c.Db = db
		c.Db = db.Debug()
	} else {
		c.Db = db
	}
}

func (c *Config) DbConn() *gorm.DB {
	return c.Db
}

func (c *Config) DbClose() {
	enigmaDb, err := c.Db.DB()
	if err != nil {
		panic(err)
	}
	defer func(enigmaDb *sql.DB) {
		err := enigmaDb.Close()
		if err != nil {
			panic(err)
		}
	}(enigmaDb)
}
func NewConfigDB() Config {
	cfg := Config{}
	cfg.initDb()
	return cfg
}
