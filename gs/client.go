// client.go
package gs

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strconv"
	"strings"
	"time"
	"tracery/protocol"
)

func init() {
	//fmt.Println("load client model")
}

type Client struct {
	Push chan []byte

	conn     *net.Conn
	lest_opt int64
	token    string
	secret   string
	handler  Handler
}

func NewClient(c *net.Conn, h Handler) (*Client, error) {
	u := &Client{
		conn:     c,
		lest_opt: time.Now().Unix(),
		Push:     make(chan []byte, 64),
		handler:  h,
	}

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

		go self.handler.handle(self, pkg)

		r.Read(make([]byte, r.Buffered()))

		fmt.Println("client>>>", strings.TrimSpace(string(body)))

	}
}

func (self *Client) send() {
	for {
		wbs := <-self.Push
		fmt.Println("write msg :", string(wbs))
		i, err := (*self.conn).Write(wbs)
		if err != nil {
			fmt.Println("send msg error :", err)
		}
		fmt.Println(i)
	}
}
