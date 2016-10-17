package convert

import . "github.com/pkg4go/assert"
import "testing"

func TestStrings(t *testing.T) {
	a := A{t}

	in1 := []string{
		"abc",
		"def",
		"ghi",
	}

	in2 := [][]byte{
		[]byte("abc"),
		[]byte("def"),
		[]byte("ghi"),
	}

	in3 := []interface{}{
		i("abc"),
		i("def"),
		i("ghi"),
	}

	out1, _ := Strings(in1)
	out2, _ := Strings(in2)
	out3, _ := Strings(in3)

	a.Equal(out1, []string{"abc", "def", "ghi"})
	a.Equal(out2, []string{"abc", "def", "ghi"})
	a.Equal(out3, []string{"abc", "def", "ghi"})
}
