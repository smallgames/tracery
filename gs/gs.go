package gs

import (
	"fmt"
	"net"
)

func init() {
	//fmt.Println("Initial gs model")
}

var (
	gs_max_onlines = 1000
	gs_max_rooms   = 10
)

type GameServer struct {
	Port   int
	Online int
	Rooms  []*GameRoom
}

func NewGS(p, c int) (*GameServer, error) {
	if gs_max_onlines > c {
		c = gs_max_onlines
	}

	grooms := make([]*GameRoom, gs_max_rooms)
	for i := 0; i < gs_max_rooms; i++ {
		grooms[i] = NewRoom(fmt.Sprintf("room_%d", i), i)
	}

	return &GameServer{Port: p, Online: c, Rooms: grooms}, nil
}

func (self *GameServer) Run() {
	tcp, err := net.ResolveTCPAddr("tcp", fmt.Sprintf(":%d", self.Port))
	if err != nil {
		panic(err)
	}

	fmt.Printf("gs start listener %s\n", tcp)
	l, err := net.ListenTCP("tcp", tcp)
	if err != nil {
		panic(err)
	}
	defer l.Close()

	for err == nil {
		if c, err := l.Accept(); err == nil {
			go NewClient(&c, NewGSHandler(self))
		} else {
			fmt.Println("gs accpet err=", err)
		}
	}
}
