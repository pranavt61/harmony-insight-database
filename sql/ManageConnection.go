package sql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"sync"
)

var DBConnection *sql.DB
var DBMutex sync.Mutex

func OpenDBConnection() {
	var err error

	DBConnection, err = sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/HarmonyDashboard")
	if err != nil {
		panic(err)
	}
}

func CloseDBConnection() {
	DBMutex.Lock()
	DBConnection.Close()
	DBMutex.Unlock()
}
