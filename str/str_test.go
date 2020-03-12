package str

import (
	"testing"
)

var cc = []byte(`xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx`)

func TestByteSlice2String(t *testing.T) {
	t.Log(ByteSlice2String(cc))
}

func Benchmark_ByteSlice2String(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = ByteSlice2String(cc)
	}
}

func Benchmark_String(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = string(cc)
	}
}
