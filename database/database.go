package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func InitDataBase() *gorm.DB {
	host := os.Getenv("HOST")
	port :=os.Getenv("DBPORT")
	user := os.Getenv("DBUSER")
	psw :=os.Getenv("DBPASSWORD")
	dbName :=os.Getenv("DBNAME")
	dsn := fmt.Sprintf("host= %s user= %s password= %s dbname= %s port= %s sslmode=disable", host, user, psw, dbName, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db
	}