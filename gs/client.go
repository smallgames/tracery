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
	fmt.Fprint(*c, "> ")

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

		datas, prefix, err := r.ReadLine()
		fmt.Println("p:", prefix)
		fmt.Println("d:", string(datas))
		if err != nil {
			fmt.Println("11111")
			fmt.Println(err)
			if err != io.EOF {
				fmt.Println("init conn declear err ", err)
				(*self.conn).Close()
				break
			}
			pkg.Mark = protocol.ERR_PKG
			r.Read(make([]byte, r.Buffered()))
			continue
		}

		body_len, err := strconv.Atoi(string(datas[:1]))
		fmt.Println("body_len:", body_len)
		if err != nil {
			fmt.Println("22222")
			pkg.Mark = protocol.ERR_PKG
			r.Read(make([]byte, r.Buffered()))
			continue
		}

		body := datas[1:body_len]
		fmt.Println("body:", string(body))
		if err != nil {
			fmt.Println("33333")
			if err != io.EOF {
				fmt.Println("init conn declear err ", err)
				(*self.conn).Close()
				break
			}
			pkg.Mark = protocol.ERR_PKG
			r.Read(make([]byte, r.Buffered()))
			continue
		}

		pkg.Head = datas[1:len(datas)]
		pkg.Body = body
		pkg.Fin = time.Now()
		pkg.Mark = protocol.FIN_PKG

		go self.handler.handle(self, pkg)

		r.Read(make([]byte, r.Buffered()))

		fmt.Println("client>>>", strings.TrimSpace(string(body)))
		fmt.Fprint((*self.conn), "> ")
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
