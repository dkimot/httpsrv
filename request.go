package httpsrv

import (
	"encoding/json"
	"net/http"
)

// ParseJSON is a helper to decode JSON bodies
func (srv *Server) ParseJSON(r *http.Request, data interface{}) error {
	decoder := json.NewDecoder(r.Body)

	return decoder.Decode(data)
}
