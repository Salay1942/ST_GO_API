package database

import "database/sql"

func ConnectDb() (db *sql.DB) {
	db, err := sql.Open("mysql", "billdb:billdb@tcp(10.0.4.53:3306)/telcomTest")
	if err != nil {
		panic(err.Error())
	}
	return db
}
