package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	got := Add(2, 2)
	want := 4

	assertCorrectMessage(t, got, want)

}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

func assertCorrectMessage(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("Got %d want %d", got, want)
	}
}
