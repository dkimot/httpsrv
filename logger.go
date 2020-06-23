package httpsrv

// Logger is an interface patterned after go-kit's logger interface
type Logger interface {
	Log(keyvals ...interface{}) error
}
