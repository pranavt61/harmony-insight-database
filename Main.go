package main

import (
	"fmt"
	"time"

	"github.com/pranavt61/harmony-insight-database/dataServer"
	"github.com/pranavt61/harmony-insight-database/nodeClient"
	"github.com/pranavt61/harmony-insight-database/sql"
)

func main() {
	sql.OpenDBConnection()
	defer sql.CloseDBConnection()

	go RoutineStartDataServer()

	go RoutineBlockTransactionCount(0)
	go RoutineBlockTransactionCount(1)
	go RoutineBlockTransactionCount(3)

	go RoutineValidators()

	for {
	}
}

func RoutineStartDataServer() {
	dataServer.StartHttpServer()
}

func RoutineBlockTransactionCount(shard_id int) {
	for {
		current_block_height := nodeClient.RequestBlockNumber(shard_id)

		start_block_height := sql.SelectMaxHeightBlockTransactionCount(shard_id) + 1
		fmt.Printf("%d - %d\n", shard_id, start_block_height)

		for height_i := start_block_height; height_i < current_block_height; height_i++ {

			// request and store
			nodeClient.RequestAndStoreBlockTransactionCount(shard_id, height_i)

			if height_i%1000 == 0 {
				fmt.Printf("SHARD %d done with %d blocks\n", shard_id, height_i)
			}
		}

		time.Sleep(60 * time.Second)
	}
}

func RoutineValidators() {
	for {
		nodeClient.RequestAndStoreValidators()

		time.Sleep(60 * time.Second)
	}
}
