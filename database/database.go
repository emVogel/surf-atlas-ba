package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
	

func InitDataBase() *gorm.DB {
	dsn := "host=172.17.0.1 user=michi password=michi1980 dbname=surfdb port=9090 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db
	}