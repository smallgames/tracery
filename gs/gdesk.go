package gs

import "fmt"

func init() {
	fmt.Println("")
}

const (
	max_player  = 4
	max_mq_size = 126
)

var ()

type GameDesk struct {
	ID      int
	Players []*Client
	MQ      chan []byte
}

func NewDesk(no int) *GameDesk {
	return &GameDesk{
		ID:      no,
		Players: make([]*Client, max_player),
		MQ:      make(chan []byte, max_mq_size),
	}
}
