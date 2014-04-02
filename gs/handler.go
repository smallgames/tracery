package gs

import (
	"bytes"
	"fmt"
	"strings"
	"tracery/protocol"
)

func init() {
	fmt.Println("load gs.handler model")
}

const ()

var ()

type Handler interface {
	handle(target *Client, p *protocol.Message)
}

//-------------------- GameServer handle
type GSHandler struct {
	Self *GameServer
}

func NewGSHandler(obj *GameServer) *GSHandler {
	return &GSHandler{
		Self: obj,
	}
}

func (self *GSHandler) handle(target *Client, p *protocol.Message) {
	switch p.Body[0] {
	case byte(protocol.LOGIN_PKG):
		fmt.Println("Login PKG")
		msg := strings.TrimSpace(string(p.Body[1:]))
		fmt.Println(msg)
		target.token = msg
		target.secret = "succeed"

		tocli := `{"rooms":[%s]}`
		gr_arr := make([]string, len(self.Self.Rooms))
		for i, v := range self.Self.Rooms {
			gr_arr[i] = fmt.Sprintf(`{"id":%d,"name":"%v","max":%d}`, v.ID, v.Name, v.Max)
		}
		tocli = fmt.Sprintf(tocli, strings.Join(gr_arr, ","))
		target.push <- bytes.NewBufferString("server>>>" + tocli).Bytes()
	case byte(protocol.GO_ROOMS_PKG):
		msg := strings.TrimSpace(string(p.Body[1:]))
		fmt.Println(msg)
	default:
		fmt.Println("Unknow package")
	}
}

//-------------------- GameRoom handle
type GRHandler struct {
	Self *GameRoom
}

func NewGRHandler(obj *GameRoom) *GRHandler {
	return &GRHandler{
		Self: obj,
	}
}

func (self *GRHandler) handle(target *Client, p *protocol.Message) {
	fmt.Println("GameRoom handler analyse body :", string(p.Body))
	target.push <- p.Body
}
