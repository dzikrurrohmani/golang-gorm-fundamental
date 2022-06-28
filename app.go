package main

import (
	"database/sql"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dbHOST := "localhost"
	dbPORT := "5432"
	dbUSER := "dzikrurrohmani"
	dbPASSWORD := "password"
	dbNAME := "golang_gorm_fundamental"
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHOST, dbUSER, dbPASSWORD, dbNAME, dbPORT)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	enigmaDb, err := db.DB()
	defer func(enigmaDb *sql.DB) {
		err := enigmaDb.Close()
		if err != nil {
			panic(err)
		}
	}(enigmaDb)

	// err = enigmaDb.Ping()
	// if err != nil {
	// 	panic(err)
	// } else {
	// 	log.Println("Connected...")
	// }

	err = db.AutoMigrate(&Customer{})
	if err != nil {
		panic(err)
	}
}
