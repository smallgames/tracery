package lib

import "testing"

func TestEnAndDe(t *testing.T) {
	i1 := 5
	bs1 := Encode(i1)
	bi1 := Decode(bs1)

	if bi1 != i1 {
		panic("not is 5")
	}

	i1 = 130
	bs1 = Encode(i1)
	bi1 = Decode(bs1)

	if bi1 != i1 {
		panic("err")
	}

	i1 = 23456
	bs1 = Encode(i1)
	bi1 = Decode(bs1)
	if bi1 != i1 {
		panic("not is 5")
	}

	i1 = 999999999999999999
	bs1 = Encode(i1)
	bi1 = Decode(bs1)
	if bi1 != i1 {
		panic("not is 5")
	}
}
