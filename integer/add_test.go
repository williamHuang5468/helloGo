package integer

import (
	"fmt"
	"testing"
)

func testAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("real is : %d, expected is : %d", sum, expected)
	}
}

func other(t *testing.T) {
	sum := Add(1, 5)
	fmt.Println(sum)
}
