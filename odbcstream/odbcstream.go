package odbcstream

import (
	"database/sql"
	"fmt"
	_"github.com/alexbrainman/odbc"
)

// A pointer to an instance of the database
var DBClient *sql.DB

// Initializing variables for data source credentials
var server, user, password, database string

type Table struct {
	name string
}

// Initializes a connection to the database of choice
func InitialiseDBConnection(driverType, dataSourceName string) error{
	dataSourceName = fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s;",
        server, user, password, database)

	db, err := sql.Open(driverType, dataSourceName)
	if err != nil {
		panic(err.Error())
	}

	// Checking if database is connected
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	DBClient = db
	return db.Close()
}

// Getting list of tables in database
func List(t *Table) {
	var tables []Table
	ta, err := DBClient.Query("SELECT * FROM sys.Table") 
		if err != nil {
			panic(err.Error())
		}
	for ta.Next() {
		var table Table
		if err := ta.Scan(&table.name); err != nil {
			fmt.Println("An error scanning for tables", err)
			return
		}
		tables = append(tables, table )
	}
}