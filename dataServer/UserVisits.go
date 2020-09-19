package dataServer

import (
	"net/http"
	"time"

	"github.com/pranavt61/harmony-insight-database/sql"
)

func handleUserVisits(w http.ResponseWriter, r *http.Request) {

	request_type := r.Method

	user_ip := ""
	forwarded_ip := r.Header.Get("X-FORWARDED-FOR")
	if forwarded_ip != "" {
		user_ip = forwarded_ip
	} else {
		user_ip = r.RemoteAddr
	}

	time := int32(time.Now().Unix())

	// store visit
	sql.InsertUserVisits(request_type, user_ip, time)
}
