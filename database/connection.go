package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() {
	dsn := "host=localhost user=postgres password=password dbname=template1 port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
    // db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    _, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

}