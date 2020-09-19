package sql

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func InsertUserVisits(request_type string, user_ip string, timestamp int32) {
	DBMutex.Lock()
	stmt, err := DBConnection.Prepare(
		`INSERT INTO UserVisits
			(
				request_type,
				user_ip,
				time
			) VALUES(?,?,?);`,
	)
	DBMutex.Unlock()
	if err != nil {
		log.Printf("SQL Statement Prepare Error: %s\n", err.Error())
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		request_type,
		user_ip,
		timestamp,
	)
	if err != nil {
		log.Printf("SQL Statement Exec Error: %s\n", err.Error())
		return
	}
}
