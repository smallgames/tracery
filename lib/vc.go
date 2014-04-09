package lib

import (
	"fmt"
)

func init() {
	fmt.Println("")
}

const ()

var ()

func Encode(i int) (buf []byte) {
	for i > 0 {
		v1 := i % 128
		buf = append(buf, byte(v1))
		if i < 128 {
			break
		}
		i /= 128
	}
	return
}

func Decode(buf []byte) int {
	n := 0
	for i := len(buf); i > 0; i-- {
		if buf[i-1]&128 != 128 {
			n = 128*n + int(buf[i-1])
		} else {
			b := buf[i-1] & 127
			n = 128*n + int(b)
			return n
		}
	}
	return n
}
