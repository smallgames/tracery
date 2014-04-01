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
	ID   int
	Max  int
	clis map[string]*Client
}

func NewRoom(n string, i int) *GameRoom {
	return &GameRoom{Name: n, ID: i, Max: gr_max_online, clis: make(map[string]*Client)}
}
