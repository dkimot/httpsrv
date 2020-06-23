package httpsrv

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

// ParseJSON is a helper to decode JSON bodies
func (srv *Server) ParseJSON(w http.ResponseWriter, r *http.Request, data interface{}) error {
	// validate the data
	if err := validate.Struct(data); err != nil {
		srv.ClientError(w, http.StatusBadRequest, nil, err)
		return err
	}

	// decode struct
	decoder := json.NewDecoder(r.Body)
	return decoder.Decode(data)
}
