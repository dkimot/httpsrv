package httpsrv

import "net/http"

// ClientError sends back a client error
func (srv *Server) ClientError(w http.ResponseWriter, status int, msg []byte, err error) {
	srv.logErr(err, status, "client")

	w.WriteHeader(status)

	if (msg == nil) {
		msg = []byte(http.StatusText(status))
	}
	w.Write(msg)
}

// InternalServerError is a helper to send a 500 response
func (srv *Server) InternalServerError(w http.ResponseWriter, err error) {
	status := http.StatusInternalServerError

	srv.logErr(err, status, "server")

	w.WriteHeader(status)

	w.Write([]byte(http.StatusText(status)))
}

// NotFoundError is a helper for 404 errors
func (srv *Server) NotFoundError(w http.ResponseWriter) {
	status := http.StatusNotFound

	srv.logErr(nil, status, "server")

	w.WriteHeader(status)

	w.Write([]byte(http.StatusText(status)))
}

func (srv *Server) logErr(err error, status int, errorType string) {
	srv.logger.Log("err", err.Error(), "status", status, "error_type", errorType)
}
