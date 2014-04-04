package gs

func init() {
	//fmt.Println("load gs.groom model")
}

const (
	gr_max_online = 200
	mq_max_size   = 1024
)

var ()

type GameRoom struct {
	Name string
	ID   int
	Max  int
	MQ   chan []byte
	clis map[string]*Client
}

func NewRoom(n string, i int) *GameRoom {
	obj := &GameRoom{Name: n,
		ID:   i,
		Max:  gr_max_online,
		MQ:   make(chan []byte, mq_max_size),
		clis: make(map[string]*Client),
	}
	return obj.builder()
}

func (self *GameRoom) builder() *GameRoom {
	go self.broadcast()
	return self
}

func (self *GameRoom) broadcast() {
	for {
		select {
		case m := <-self.MQ:
			for _, v := range self.clis {
				v.Push <- m
			}
		}
	}
}
