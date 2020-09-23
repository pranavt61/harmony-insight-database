package sql

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type RowBlockGasUsed struct {
	Shard_id     int
	Block_height int
	Gas_used     int
}

func InsertBlockGasUsed(shard_id int, block_height int, gas_used int) {

	DBMutex.Lock()
	stmt, err := DBConnection.Prepare(
		`INSERT INTO BlockGasUsed
			(
				shard_id,
				block_height,
				gas_used
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
		gas_used,
	)
	if err != nil {
		log.Printf("SQL Statement Exec Error: %s\n", err.Error())
		return
	}
}

func SelectBlockGasUsed(min_block_height int) []RowBlockGasUsed {
	DBMutex.Lock()
	ret, err := DBConnection.Query(
		`SELECT 
			*
		FROM BlockGasUsed
		WHERE block_height > ?;`,
		min_block_height,
	)
	DBMutex.Unlock()
	if err != nil {
		log.Printf("SQL Statement Prepare Error: %s\n", err.Error())
		return nil
	}

	rows := make([]RowBlockGasUsed, 0)
	row_buffer := RowBlockGasUsed{}
	for ret.Next() {
		ret.Scan(
			&(row_buffer.Shard_id),
			&(row_buffer.Block_height),
			&(row_buffer.Gas_used),
		)

		row := RowBlockGasUsed{
			row_buffer.Shard_id,
			row_buffer.Block_height,
			row_buffer.Gas_used,
		}
		rows = append(rows, row)
	}

	return rows
}

func SelectMaxHeightBlockGasUsed(shard_id int) int {
	DBMutex.Lock()
	ret, err := DBConnection.Query(
		`SELECT 
			MAX(block_height)
		FROM BlockGasUsed
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
