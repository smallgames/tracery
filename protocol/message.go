package protocol

import (
	"fmt"
	"time"
)

func init() {
	fmt.Println("load msg model")
}

const (
	NEW_PKG = iota
	TCZ_PKG
	FIN_PKG
	INV_PKG
	ERR_PKG
)

type Message struct {
	Mark int
	Fin  time.Time
	Head []byte
	Body []byte
}

func (self *Message) Test() []byte {
	return self.Body
}
