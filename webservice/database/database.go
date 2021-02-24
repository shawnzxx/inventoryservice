package database

import (
	"database/sql"
	"log"
	"time"
)

//DbConn - export connected DbConn object
var DbConn *sql.DB

//SetupDatabase - connect to the db
func SetupDatabase() {
	var err error
	DbConn, err = sql.Open("mysql", "root:Passw0rd123!@tcp(127.0.0.1:3306)/inventorydb")
	if err != nil {
		log.Fatal(err)
	}
	DbConn.SetMaxOpenConns(3)
	DbConn.SetMaxIdleConns(3)
	DbConn.SetConnMaxLifetime(60 * time.Second)
}
