// client.go
package gs

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strings"
	"time"
)

func init() {
	//fmt.Println("load client model")
}

type Client struct {
	Push chan []byte

	conn     net.Conn
	live     bool
	end      bool
	lest_opt int64
	token    string
	secret   string
	handler  Handler
}

func NewClient(c net.Conn, h Handler) (*Client, error) {
	u := Client{
		conn:     c,
		live:     true,
		end:      false,
		lest_opt: time.Now().Unix(),
		Push:     make(chan []byte, 64),
		handler:  h,
	}
	fmt.Fprint(c, "> ")

	go u.receive()
	go u.send()
	return &u, nil
}

func (self Client) receive() {
	r := bufio.NewReader(self.conn)

	for self.live {
		datas, _, err := r.ReadLine()
		if err != nil {
			fmt.Println(err)
			if err != io.EOF {
				if !self.end {
					self.Close()
				}
				break
			}
			if !self.end {
				self.Close()
			}
			break
			continue
		}

		fmt.Println("body:", string(datas))
		if err != nil {
			fmt.Println("33333")
			if err != io.EOF {
				if !self.end {
					self.conn.Close()
				}
				break
			}
			if !self.end {
				self.conn.Close()
			}
			break
			continue
		}

		go self.handler.handle(self, datas)

		//r.Read(make([]byte, r.Buffered()))

		fmt.Println("client>>>", strings.TrimSpace(string(datas)))
		fmt.Fprint(self.conn, "server> "+strings.TrimSpace(string(datas))+"\n> ")
	}
}

func (self Client) send() {
	for self.live {
		wbs := <-self.Push
		fmt.Println("write msg :", string(wbs))
		i, err := self.conn.Write(wbs)
		if err != nil {
			fmt.Println("send msg error :", err)
		}
		fmt.Println(i)
	}
}

func (self Client) Close() {
	self.live = false
	if !self.end {
		self.conn.Close()
	}
}
