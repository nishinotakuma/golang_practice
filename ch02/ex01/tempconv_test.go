package tempconv

import (
	"testing"
)

func TestTempConv(t *testing.T) {
	type Test struct {
		c Celsius
		k Kelvin
	}

	test := Test{
		c: 30,
		k: 303.15,
	}

	actual := KToC(test.k)
	expected := test.c
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
