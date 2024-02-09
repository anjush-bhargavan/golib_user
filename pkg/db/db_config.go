package db

import (
	"log"
	"os"

	dom "github.com/anjush-bhargavan/golib_user/pkg/DOM"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	dsn, ok := os.LookupEnv("DB_Config")
	if !ok {
		log.Fatal("error loading database env")
	}

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("error conncting to databse: %v", err.Error())
	}
	DB.AutoMigrate(&dom.User{},
		&dom.Membership{})

	return DB
}
