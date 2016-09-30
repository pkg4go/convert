package convert

import . "github.com/pkg4go/assert"
import "testing"

// convert to string

func TestConvertStringToString(t *testing.T) {
	a := A{t}
	a.Equal(String("123456"), "123456")
}

func TestConvertIntToString(t *testing.T) {
	a := A{t}
	var in int = 10086
	a.Equal(String(in), "10086")
}

func TestConvertInt32ToString(t *testing.T) {
	a := A{t}
	var in int32 = 12345
	a.Equal(String(in), "12345")
}

func TestConvertInt64ToString(t *testing.T) {
	a := A{t}
	var in int64 = 1258096
	a.Equal(String(in), "1258096")
}

func TestConvertFloat32ToString(t *testing.T) {
	a := A{t}
	var in float32 = 123.321
	a.Equal(String(in, 3), "123.321")
}

func TestConvertFloat64ToString(t *testing.T) {
	a := A{t}
	var in float64 = 1234567.7654321
	a.Equal(String(in, 7), "1234567.7654321")
}

func TestConvertArrayToString(t *testing.T) {
	a := A{t}
	a.Equal(String(i("hello")), "hello")
	a.Equal(String([]byte("hello")), "hello")
	a.Equal(String(i([]byte("hello"))), "hello")
	a.Equal(String([]byte("world 😂")), "world 😂")
}

// convert to int

func TestConvertStringToInt(t *testing.T) {
	a := A{t}
	in := "123"
	out, _ := Int(in)
	a.Equal(out, int(123))
}

func TestConvertIntToInt(t *testing.T) {
	a := A{t}
	var in int = 10086
	out, _ := Int(in)
	a.Equal(out, int(in))
}

func TestConvertInt32ToInt(t *testing.T) {
	a := A{t}
	var in int32 = 12345
	out, _ := Int(in)
	a.Equal(out, int(in))
}

func TestConvertInt64ToInt(t *testing.T) {
	a := A{t}
	var in int64 = 1258096
	out, _ := Int(in)
	a.Equal(out, int(in))
}

func TestConvertFloat32ToInt(t *testing.T) {
	a := A{t}
	var in float32 = 123.321
	out, _ := Int(in)
	a.Equal(out, int(123))
}

func TestConvertFloat64ToInt(t *testing.T) {
	a := A{t}
	var in float64 = 1234567.7654321
	out, _ := Int(in)
	a.Equal(out, int(1234567))
}

// convert to int32

func TestConvertStringToInt32(t *testing.T) {
	a := A{t}
	in := "123"
	out, _ := Int32(in)
	a.Equal(out, int32(123))
}

func TestConvertIntToInt32(t *testing.T) {
	a := A{t}
	var in int = 10086
	out, _ := Int32(in)
	a.Equal(out, int32(in))
}

func TestConvertInt32ToInt32(t *testing.T) {
	a := A{t}
	var in int32 = 12345
	out, _ := Int32(in)
	a.Equal(out, int32(in))
}

func TestConvertInt64ToInt32(t *testing.T) {
	a := A{t}
	var in int64 = 1258096
	out, _ := Int32(in)
	a.Equal(out, int32(in))
}

func TestConvertFloat32ToInt32(t *testing.T) {
	a := A{t}
	var in float32 = 123.321
	out, _ := Int32(in)
	a.Equal(out, int32(123))
}

func TestConvertFloat64ToInt32(t *testing.T) {
	a := A{t}
	var in float64 = 1234567.7654321
	out, _ := Int32(in)
	a.Equal(out, int32(1234567))
}

// convert to int64

func TestConvertStringToInt64(t *testing.T) {
	a := A{t}
	in := "123"
	out, _ := Int64(in)
	a.Equal(out, int64(123))
}

func TestConvertIntToInt64(t *testing.T) {
	a := A{t}
	var in int = 10086
	out, _ := Int64(in)
	a.Equal(out, int64(in))
}

func TestConvertInt32ToInt64(t *testing.T) {
	a := A{t}
	var in int32 = 12345
	out, _ := Int64(in)
	a.Equal(out, int64(in))
}

func TestConvertInt64ToInt64(t *testing.T) {
	a := A{t}
	var in int64 = 1258096
	out, _ := Int64(in)
	a.Equal(out, int64(in))
}

func TestConvertFloat32ToInt64(t *testing.T) {
	a := A{t}
	var in float32 = 123.321
	out, _ := Int64(in)
	a.Equal(out, int64(123))
}

func TestConvertFloat64ToInt64(t *testing.T) {
	a := A{t}
	var in float64 = 1234567.7654321
	out, _ := Int64(in)
	a.Equal(out, int64(1234567))
}

// convert to float32

func TestConvertStringToFloat32(t *testing.T) {
	a := A{t}
	in := "123.456"
	out, _ := Float32(in)
	a.Equal(out, float32(123.456))
}

func TestConvertIntToFloat32(t *testing.T) {
	a := A{t}
	var in int = 10086
	out, _ := Float32(in)
	a.Equal(out, float32(in))
}

func TestConvertInt32ToFloat32(t *testing.T) {
	a := A{t}
	var in int32 = 12345
	out, _ := Float32(in)
	a.Equal(out, float32(in))
}

func TestConvertInt64ToFloat32(t *testing.T) {
	a := A{t}
	var in int64 = 1258096
	out, _ := Float32(in)
	a.Equal(out, float32(in))
}

func TestConvertFloat32ToFloat32(t *testing.T) {
	a := A{t}
	var in float32 = 123.321
	out, _ := Float32(in)
	a.Equal(out, float32(in))
}

func TestConvertFloat64ToFloat32(t *testing.T) {
	a := A{t}
	var in float64 = 1234567.7654321
	out, _ := Float32(in)
	a.Equal(out, float32(in))
}

// convert to float64

func TestConvertStringToFloat64(t *testing.T) {
	a := A{t}
	in := "123.456"
	out, _ := Float64(in)
	a.Equal(out, float64(123.456))
}

func TestConvertIntToFloat64(t *testing.T) {
	a := A{t}
	var in int = 10086
	out, _ := Float64(in)
	a.Equal(out, float64(in))
}

func TestConvertInt32ToFloat64(t *testing.T) {
	a := A{t}
	var in int32 = 12345
	out, _ := Float64(in)
	a.Equal(out, float64(in))
}

func TestConvertInt64ToFloat64(t *testing.T) {
	a := A{t}
	var in int64 = 1258096
	out, _ := Float64(in)
	a.Equal(out, float64(in))
}

func TestConvertFloat32ToFloat64(t *testing.T) {
	a := A{t}
	var in float32 = 123.321
	out, _ := Float64(in)
	a.Equal(out, float64(in))
}

func TestConvertFloat64ToFloat64(t *testing.T) {
	a := A{t}
	var in float64 = 1234567.7654321
	out, _ := Float64(in)
	a.Equal(out, float64(in))
}

// utils
func i(v interface{}) interface{} {
	return v
}
