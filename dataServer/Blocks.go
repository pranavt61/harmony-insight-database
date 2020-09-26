package dataServer

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/pranavt61/harmony-insight-database/nodeClient"
)

func handleBlockHashByNumber(w http.ResponseWriter, r *http.Request) {

	block_height := 0
	var err error

	block_height_string := r.URL.Query().Get("block_height")
	if block_height_string != "" {
		block_height, err = strconv.Atoi(block_height_string)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error parsing GET parameters 'block_height'")

			return
		}
	}

	hash := nodeClient.RequestBlockHashByNumber(block_height)

	fmt.Fprintf(w, `{"hash": "%s"}`, hash)
}
