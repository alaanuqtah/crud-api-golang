package initializers

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB //create a public var of type gorm.DB which means the local db for this program

func ConnectToDB() {
	var err error
	dsn := os.Getenv("DB_URL") //get env var url which contains the config of the connection to the remote db (postgress)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("failed to connect to database")
	}
}
