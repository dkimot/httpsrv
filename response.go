package httpsrv

import (
	"encoding/json"
	"net/http"
)

// RespondJSON is a helper to encode and send data as a json response
func (srv *Server) RespondJSON(w http.ResponseWriter, status int, data interface{}) {
	body, err := json.Marshal(data)
	if err != nil {
		srv.InternalServerError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_, err = w.Write(body)
	if err != nil {
		srv.logErr(err, 0, "server")
	}
}
