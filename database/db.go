package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectWithDB() {
	con := "user=admin dbname=personalities password=Passw0rd host=localhost sslmode=disable"
	DB, err = gorm.Open(postgres.Open(con))
	if err != nil {
		log.Panic("db connection error: ", err)
	}
}
