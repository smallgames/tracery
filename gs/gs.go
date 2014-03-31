package gs

import (
	"bufio"
	//"bytes"
	"fmt"
	"io"
	"net"
	"strconv"
	"strings"
	"time"
	"tracery/lib"
	"tracery/protocol"
)

func init() {
	fmt.Println("Initial gs model")
}

var (
	def_capacity = 200
)

type GameServer struct {
	Port     int
	Capacity int
	SysInfo  *lib.Task
}

func NewGS(p int, c int, t *lib.Task) (*GameServer, error) {
	if def_capacity > c {
		c = def_capacity
	}
	return &GameServer{Port: p, Capacity: c, SysInfo: t}, nil
}

func (self *GameServer) Run() {
	tcp, err := net.ResolveTCPAddr("tcp", fmt.Sprintf(":%d", self.Port))
	if err != nil {
		fmt.Errorf("port formatter err", err)
		return
	}

	fmt.Printf("gs start listener %s\n", tcp)
	l, err := net.ListenTCP("tcp", tcp)
	if err != nil {
		fmt.Errorf("port formatter err", err)
		return
	}
	defer l.Close()

	for err == nil {
		if c, err := l.Accept(); err == nil {
			go init_conn(&c)
		} else {
			fmt.Println("gs accpet err=", err)
		}

	}
}

type Client struct {
	conn     *net.Conn
	lest_opt int64
	token    string
	secret   string
	push     chan []byte
}

func init_conn(c *net.Conn) (*Client, error) {
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
			continue
		}

		body_len, err := strconv.Atoi(string(head[:len(head)-1]))
		if err != nil {
			pkg.Mark = protocol.ERR_PKG
			continue
		}

		body := make([]byte, body_len+2)
		i, err := r.Read(body)
		if err != nil {
			if err != io.EOF {
				fmt.Println("init conn declear err ", err)
			}
			pkg.Mark = protocol.ERR_PKG
			break
		}

		pkg.Head = head[:len(head)-1]
		pkg.Body = body
		pkg.Fin = time.Now()

		fmt.Println(strings.TrimSpace(string(body)))
		fmt.Println(i)

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
	msg := strings.TrimSpace(string(p.Body))
	fmt.Println(msg)
	self.push <- p.Test()
}
