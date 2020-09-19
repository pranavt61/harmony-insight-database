package dataServer

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pranavt61/harmony-insight-database/sql"
)

func handleValidatorByAddress(w http.ResponseWriter, r *http.Request) {

	address := r.URL.Query().Get("address")
	if address == "" {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error parsing GET parameters\n")
	}

	row := sql.SelectValidatorByAddress(address)

	row_bytes, err := json.Marshal(row)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error parsing rows\n")
	}

	w.Write(row_bytes)
}
