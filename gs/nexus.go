package gs

func init() {
	//fmt.Println("load nexus model")
}

type Nexus struct {
	sessions map[string]*Client
}

func NewNexus() *Nexus {
	return &Nexus{sessions: make(map[string]*Client)}
}

func (self *Nexus) Add(n string, c *Client) {
	self.sessions[n] = c
}
