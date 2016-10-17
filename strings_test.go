package convert

import . "github.com/pkg4go/assert"
import "testing"

func TestStrings(t *testing.T) {
	a := A{t}

	in := [][]byte{
		[]byte("abc"),
		[]byte("def"),
		[]byte("ghi"),
	}

	out, _ := Strings(in)

	a.Equal(out, []string{"abc", "def", "ghi"})
}
