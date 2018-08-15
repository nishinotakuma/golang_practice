package reverse

import (
	"testing"
)

func TestRevUTF8(t *testing.T) {
	s := []byte("いあGFEdcba")
	got := string(revUTF8(s))
	want := "adcdEFGあい"
	if got != want {
		t.Errorf("got %v, want %v", string(got), want)
	}
}
