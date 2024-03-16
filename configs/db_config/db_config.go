package db_config

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Table struct {      
	Tags	string
	Customers	string
}

func Connect() (db *gorm.DB, tb *Table) {
	GO_SUPABASE_CONNECTION_STRING := os.Getenv("GO_SUPABASE_CONNECTION_STRING")
	// 
	pg := postgres.Open(GO_SUPABASE_CONNECTION_STRING)
	db, err := gorm.Open(pg, &gorm.Config{})
	
	if err != nil {
		panic(err)
	}

	return db, &Table{
		Tags: "tags",
		Customers: "customers",
	}
}