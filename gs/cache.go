package gs

import (
	"errors"
	"fmt"
	"reflect"
	//"sync"
)

func init() {
	fmt.Println("load memcache model")
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
	Get() reflect.Value
	Len() int
	Append(Memcache) error
}

// byte of memcache begin
type BytesMem struct {
	//lock sync.RWMutex
	data []byte
	head *BytesMem
	pre  *BytesMem
	next *BytesMem
}

func (self *BytesMem) Get() reflect.Value {
	return reflect.ValueOf(self.data)
}
func (self *BytesMem) Len() int {
	return len(self.data)
}
func (self *BytesMem) Append(m Memcache) error {
	if p, ok := m.(*BytesMem); ok {
		if self.next == nil {
			self.next = p
			p.pre = self
			p.head = self.head
		} else {
			self.next.Append(p)
		}
		return nil
	}
	return err_not_same_type

}

// share funcs begin
func NewMem(mt int) (Memcache, error) {
	switch mt {
	case BYTES_MEM:
		return &BytesMem{data: make([]byte, 1, 2)}, nil
	default:
		return nil, errors.New(fmt.Sprintf(err_unknow_type_str, mt))
	}
	return nil, err_not_same_type
}
