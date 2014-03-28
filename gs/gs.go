package gs

import (
	"fmt"
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
	fmt.Println("gs start gogogo")
	defer fmt.Println("gs end")
	for {
		select {
		case <-time.After(time.Second * 5):
			fmt.Println("hahah")
		}
	}
}
