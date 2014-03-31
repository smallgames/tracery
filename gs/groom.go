package gs

import (
	"fmt"
)

func init() {
	fmt.Println("load gs.groom model")
}

const (
	gr_max_online = 200
)

var ()

type GameRoom struct {
	Name string
	Id   int
	Max  int
	clis map[string]*Client
}

func NewRoom(n string, id int) *GameRoom {
	return &GameRoom{Name: n, Max: gr_max_online, clis: make([string]*Client)}
}
