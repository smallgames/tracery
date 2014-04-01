// client.go
package gs

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net"
	"strconv"
	"strings"
	"time"
	"tracery/protocol"
)

func init() {
	fmt.Println("load client model")
}

type Client struct {
	conn     *net.Conn
	lest_opt int64
	token    string
	secret   string
	push     chan []byte
}

func NewClient(c *net.Conn) (*Client, error) {
	u := &Client{conn: c, lest_opt: time.Now().Unix(), push: make(chan []byte, 64)}
	go u.receive()
	go u.send()
	return u, nil
}

func (self *Client) receive() {
	pkg := &protocol.Message{Mark: protocol.NEW_PKG}
	r := bufio.NewReader((*self.conn))

	for {
		if pkg.Mark > 1 {
			fmt.Println("receive new pkg")
			pkg = &protocol.Message{Mark: protocol.NEW_PKG}
		}

		head, err := r.ReadSlice(byte('|'))
		if err != nil {
			if err != io.EOF {
				fmt.Println("init conn declear err ", err)
			}
			pkg.Mark = protocol.ERR_PKG
			r.Read(make([]byte, r.Buffered()))
			continue
		}

		body_len, err := strconv.Atoi(string(head[:len(head)-1]))
		if err != nil {
			pkg.Mark = protocol.ERR_PKG
			r.Read(make([]byte, r.Buffered()))
			continue
		}

		body := make([]byte, body_len)
		_, err = r.Read(body)
		if err != nil {
			if err != io.EOF {
				fmt.Println("init conn declear err ", err)
			}
			pkg.Mark = protocol.ERR_PKG
			r.Read(make([]byte, r.Buffered()))
			continue
		}

		pkg.Head = head[:len(head)-1]
		pkg.Body = body
		pkg.Fin = time.Now()
		pkg.Mark = protocol.FIN_PKG

		go self.handle(pkg)

		r.Read(make([]byte, r.Buffered()))

		fmt.Println("client>>>", strings.TrimSpace(string(body)))

	}
}

func (self *Client) send() {
	for {
		wbs := <-self.push
		fmt.Println("write msg :", string(wbs))
		i, err := (*self.conn).Write(wbs)
		if err != nil {
			fmt.Println("send msg error :", err)
		}
		fmt.Println(i)
	}
}

func (self *Client) handle(p *protocol.Message) {
	//msg := strings.TrimSpace(string(p.Body))
	//fmt.Println(msg)
	//fmt.Println(p.Body[0])
	//fmt.Println(byte(protocol.LOGIN_PKG))
	switch p.Body[0] {
	case byte(protocol.LOGIN_PKG):
		msg := strings.TrimSpace(string(p.Body[1:]))
		fmt.Println(msg)
		self.token = msg
		self.secret = "succeed"
		self.push <- bytes.NewBufferString("请选择大厅").Bytes()
	case byte(protocol.GO_ROOMS_PKG):
		msg := strings.TrimSpace(string(p.Body[1:]))
		fmt.Println(msg)
	default:
		fmt.Println("Unknow package")
	}
}
