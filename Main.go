package main

import (
	"fmt"
	"github.com/pranavt61/harmony-insight-database/nodeClient"
	"github.com/pranavt61/harmony-insight-database/sql"
)

func main() {
	sql.OpenDBConnection()
	defer sql.CloseDBConnection()

	go RoutineBlockTransactionCount(0)
	go RoutineBlockTransactionCount(1)
	go RoutineBlockTransactionCount(3)

	for {
	}
}

func RoutineBlockTransactionCount(shard_id int) {

	max_block_height := nodeClient.RequestBlockNumber(shard_id)

	for height_i := 0; height_i < max_block_height; height_i++ {
		nodeClient.RequestBlockTransactionCount(shard_id, height_i)

		if height_i%1000 == 0 {
			fmt.Printf("SHARD %d done with %d blocks\n", shard_id, height_i)
		}
	}
}
