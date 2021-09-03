// This package have the database related all files
package database

// Import require package for mysql db
import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func dbConn() (db *sql.DB) {  //DATABASE Connection
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "123456"
	dbName := "sugarboxdb"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
        	panic(err.Error())
	}
	return db
}
