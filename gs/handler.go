package gs

import "fmt"
import "tracery/protocol"

func init() {

}

const ()

var ()

type Handler interface {
	handle(target Client, b []byte)
}

type ClientHandler struct {
}

func NewHandler() ClientHandler {
	return ClientHandler{}
}

func (self ClientHandler) handle(target Client, b []byte) {
	p := protocol.Reader(b)
	t, err := p.ReadU16()
	if err != nil {
		fmt.Errorf("发现未知报文")
	}
	fmt.Println(t)
	switch t {
	case protocol.QUIT_PKG:
		fmt.Println("QUIT PKG")
		target.Close()

	case protocol.LOGIN_PKG:
		fmt.Println("Login PKG")
		msg, err := p.ReadString()
		if err != nil {
			fmt.Errorf("登陆GS出错", err)
		}
		fmt.Println(msg)
		target.token = msg
		target.secret = "succeed"
		target.Push <- []byte("LOGIN SUCCEED")
		fmt.Println("LOGIN SUCCEED")
	case protocol.GO_ROOMS_PKG:
		msg, err := p.ReadString()
		if err != nil {
			fmt.Errorf("登陆ROOT出错", err)
		}
		fmt.Println("???", msg)
	default:
		fmt.Println("Unknow package")
		target.Close()
	}
}
