package gs

import (
	"errors"
	"fmt"
	//"reflect"
	//"sync"
)

func init() {
	//fmt.Println("load memcache model")
}

const (
	BYTES_MEM = iota

	DEF_BYTES_LEN = 1024
	DEF_BYTES_CAP = 4096
)

var (
	err_not_same_type = errors.New("Not same type memcache")

	err_unknow_type_str = "Unknow type %d"
)

type Memcache interface {
	Set([]byte) error
	Get() *[]byte
	Len() int
	Append(Memcache) error
}

// byte of memcache begin
type BytesMem struct {
	//lock sync.RWMutex
	Data *[]byte
	Head *BytesMem
	Pre  *BytesMem
	Next *BytesMem
}

// implement Memcache
func (self *BytesMem) Set(v []byte) error {
	self.Data = &v
	return nil
}
func (self *BytesMem) Get() *[]byte {
	return self.Data
}
func (self *BytesMem) Len() int {
	return len(*self.Data)
}
func (self *BytesMem) Append(m Memcache) error {
	if p, ok := m.(*BytesMem); ok {
		if self.Next == nil {
			self.Next = p
			p.Pre = self
			p.Head = self.Head
		} else {
			self.Next.Append(p)
		}
		return nil
	}
	return err_not_same_type

}

func NewMem(mt int) (Memcache, error) {
	switch mt {
	case BYTES_MEM:
		return &BytesMem{}, nil
	default:
		return nil, errors.New(fmt.Sprintf(err_unknow_type_str, mt))
	}
	return nil, err_not_same_type
}

//func NewMem(n string, mt int) (Memcache, error) {
//	switch mt {
//	case BYTES_MEM:
//		this := &BytesMem{Name: n, Data: make([]byte, 1, 2)}
//		this.Head = this
//		return this, nil
//	default:
//		return nil, errors.New(fmt.Sprintf(err_unknow_type_str, mt))
//	}
//	return nil, err_not_same_type
//}
