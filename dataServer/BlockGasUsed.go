package dataServer

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/pranavt61/harmony-insight-database/sql"
)

func handleBlockGasUsed(w http.ResponseWriter, r *http.Request) {

	min_block_height := 0
	var err error

	min_block_height_string := r.URL.Query().Get("min_block_height")
	if min_block_height_string != "" {
		min_block_height, err = strconv.Atoi(min_block_height_string)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error parsing GET parameters 'min_block_height'")
			return
		}
	}

	rows := sql.SelectBlockGasUsed(min_block_height)

	rows_bytes, err := json.Marshal(rows)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error parsing rows\n")
		return
	}

	w.Write(rows_bytes)
}

func handleMaxBlockHeightBlockGasUsed(w http.ResponseWriter, r *http.Request) {

	height := sql.SelectMaxHeightBlockGasUsed(0)

	fmt.Fprintf(w, `{"height": "%d"}`, height)
}
