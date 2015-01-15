package protocol

import "time"

func init() {
}

const (
	NEW_PKG = iota
	TCZ_PKG
	FIN_PKG
	INV_PKG
	ERR_PKG
)

const (
	QUIT_PKG     = 12336 //00
	LOGIN_PKG    = 12337 //01
	GO_ROOMS_PKG = 12338 //02
)

type Message struct {
	Mark int
	Fin  time.Time
	Body []byte
}
