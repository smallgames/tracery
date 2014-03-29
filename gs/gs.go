package gs

import (
	//"bufio"
	//"bytes"
	"fmt"
	"io"
	"net"
	//"strings"
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

	fmt.Printf("gs start listener %s", tcp)
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
}

func init_conn(c *net.Conn) (*Client, error) {
	u := &Client{conn: c, lest_opt: time.Now().Unix(), token: "", secret: ""}

	var (
		n, o Memcache
	)

	n, err := NewMem(BYTES_MEM)
	if err != nil {
		return nil, (*c).Close()
	}
	for {
		d := make([]byte, 1024, 4096)
		i, err := (*c).Read(d)
		if err != nil {
			if err != io.EOF {
				fmt.Println("init conn declear err ", err)
			}
			break
		}
		fmt.Println(i)
		n.Set(d[:i])
		if o != nil {
			o.Append(n)
		}
		o = n
		n, err = NewMem(BYTES_MEM)
		if err != nil {
			fmt.Println("create bm error", err)
			break
		}
	}
	return u, nil
}
