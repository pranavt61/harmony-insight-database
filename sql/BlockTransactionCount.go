package sql

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func InsertBlockTransactionCount(shard_id int, block_height int, tx_count int) {

	DBMutex.Lock()
	stmt, err := DBConnection.Prepare(
		`INSERT INTO BlockTransactionCount
			(
				shard_id,
				block_height,
				tx_count
			) VALUES(?,?,?);`,
	)
	DBMutex.Unlock()
	if err != nil {
		log.Printf("SQL Statement Prepare Error: %s\n", err.Error())
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		shard_id,
		block_height,
		tx_count,
	)
	if err != nil {
		log.Printf("SQL Statement Exec Error: %s\n", err.Error())
		return
	}
}
