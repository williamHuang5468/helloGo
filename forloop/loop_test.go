package forloop

import (
	"testing"
)

func TestLoop(t *testing.T) {
	loop := Repeat("a")
	expected := "aaaaa"

	if loop != expected {
		t.Errorf("loop is : %s, expected is : %s", loop, expected)
	}
}

/*
How to execute  the performance test?
$ go test -bench=.
*/
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
