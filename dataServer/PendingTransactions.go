package dataServer

import (
  "fmt"
  "strconv"
	"net/http"

	"github.com/pranavt61/harmony-insight-database/nodeClient"
)

func handlePendingTransactions(w http.ResponseWriter, r *http.Request) {

  shard_id := 0
  var err error

	shard_id_string := r.URL.Query().Get("shard_id")
	if shard_id_string != "" {
		shard_id, err = strconv.Atoi(shard_id_string)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error parsing GET parameters 'shard_id'")

			return
		}
	}

  pending_txs := nodeClient.RequestPendingTransactions(shard_id)

  fmt.Fprintf(w, `{"pending_txs": %s}`, pending_txs)
}
