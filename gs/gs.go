package gs

import (
	"fmt"
	"net"
	"tracery/lib"
)

func init() {
	fmt.Println("Initial gs model")
}

var (
	gs_max_onlines = 1000
	gs_max_rooms   = 10
)

type GameServer struct {
	SysInfo *lib.Task
	Port    int
	Online  int
	Rooms   []*GameRoom

	//clis map[string]*Client
}

func NewGS(p int, c int, t *lib.Task) (*GameServer, error) {
	if gs_max_onlines > c {
		c = gs_max_onlines
	}

	grooms := make([]*GameRoom, gs_max_rooms)
	for i := 0; i < gs_max_rooms; i++ {
		grooms[i] = NewRoom(fmt.Sprintf("room_", i), i)
	}

	return &GameServer{Port: p, Online: c, SysInfo: t, Rooms: grooms}, nil
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
			go NewClient(&c)
		} else {
			fmt.Println("gs accpet err=", err)
		}
	}
}
