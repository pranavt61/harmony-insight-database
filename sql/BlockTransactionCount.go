package sql

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type RowBlockTransactionCount struct {
	Shard_id     int
	Block_height int
	Tx_count     int
}

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

func SelectBlockTransactionCount(min_block_height int) []RowBlockTransactionCount {
	DBMutex.Lock()
	ret, err := DBConnection.Query(
		`SELECT 
			*
		FROM BlockTransactionCount
		WHERE block_height > ?;`,
		min_block_height,
	)
	DBMutex.Unlock()
	if err != nil {
		log.Printf("SQL Statement Prepare Error: %s\n", err.Error())
		return nil
	}

	rows := make([]RowBlockTransactionCount, 0)
	row_buffer := RowBlockTransactionCount{}
	for ret.Next() {
		ret.Scan(
			&(row_buffer.Shard_id),
			&(row_buffer.Block_height),
			&(row_buffer.Tx_count),
		)

		row := RowBlockTransactionCount{
			row_buffer.Shard_id,
			row_buffer.Block_height,
			row_buffer.Tx_count,
		}
		rows = append(rows, row)
	}

	return rows
}

func SelectMaxHeightBlockTransactionCount(shard_id int) int {
	DBMutex.Lock()
	ret, err := DBConnection.Query(
		`SELECT 
			MAX(block_height)
		FROM BlockTransactionCount
		WHERE shard_id = ?;`,
		shard_id,
	)
	DBMutex.Unlock()
	if err != nil {
		log.Printf("SQL Statement Prepare Error: %s\n", err.Error())
		return -1
	}

	height := 0
	for ret.Next() {
		ret.Scan(
			&height,
		)
	}

	return height
}
