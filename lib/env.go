package lib

import (
	"fmt"
	"io"
	"os"
	"time"
)

func init() {
	fmt.Println("Loaded env modle")
}

type Conf struct {
	Last_Opt time
	File     string
	Stroes   map[string]*_item
}

type _item struct {
	key string
	val *interface{}
}

func NewConf(fs string) (*Conf, error) {
	f, err := os.OpenFile(f, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}

	stores, err := _init(f)
	if err != nil {
		return nil, err
	}

	return &Conf{Last_Opt: time.Now(), File: fs, Stroes: stores}
}

func _init(f *os.File) (map[string]*_item, error) {
	val := make(map[string]*_item)
	return val, nil
}
