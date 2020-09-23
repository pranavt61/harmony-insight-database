package dataServer

import (
	"net/http"

	"github.com/gorilla/mux"
)

// HTTP server wrapper
type Server struct {
	router *mux.Router
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// add headers
	w.Header().Set("Access-Control-Allow-Origin", "*")

	s.router.ServeHTTP(w, r)
}

func StartHttpServer() {

	server := &Server{}
	router := mux.NewRouter()

	// BlockTransactionCount
	router.HandleFunc("/block_transaction_count", handleBlockTransactionCount)
	router.HandleFunc("/max_block_height_block_transaction_count", handleMaxBlockHeightBlockTransactionCount)

	// BlockGasUsed
	router.HandleFunc("/block_gas_used", handleBlockGasUsed)
	router.HandleFunc("/max_block_height_block_gas_used", handleMaxBlockHeightBlockGasUsed)

	// UserVisits
	router.HandleFunc("/user_visits", handleUserVisits)

	// Validators
	router.HandleFunc("/validator_by_address", handleValidatorByAddress)

	server.router = router

	http.Handle("/", server)
	http.ListenAndServe(":8081", nil)
}
