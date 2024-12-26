package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func InitDataBase() *gorm.DB {
	host := "172.17.0.1" //os.Getenv("HOST")
	port := "9090" //os.Getenv("DBPORT")
	user := "michi" //os.Getenv("DBUSER")
	psw := "michi1980" //os.Getenv("DBPASSWORD")
	dbName := "surfdb" // os.Getenv("DBNAME")
	dsn := fmt.Sprintf("host= %s user= %s password= %s dbname= %s port= %s sslmode=disable", host, user, psw, dbName, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db
	}