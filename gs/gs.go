package gs

import (
	"fmt"
	"net"
	"time"
	"tracery/lib"
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
	//fmt.Println("gs start gogogo")
	//defer fmt.Println("gs end")

	tcp, err := net.ResolveTCPAddr("tcp", fmt.Sprintf(":%d", self.Port))
	if err != nil {
		fmt.Errorf("port formatter err", err)
		return
	}

	l, err := net.ListenTCP("tcp", tcp)
	if err != nil {
		fmt.Errorf("port formatter err", err)
		return
	}
	defer l.Close()

	for c, err := l.Accept(); err == nil; {
		go init_user(&c)
	}
}

type User struct {
	conn     *net.Conn
	lest_opt int64
	token    string
	secret   string
}

func init_user(c *net.Conn) (*User, error) {
	u := &User{conn: c, lest_opt: time.Now().Unix(), token: "", secret: ""}
	for {
		c.Read()
	}
	return u, nil
}
