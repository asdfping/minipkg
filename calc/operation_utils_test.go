package calc

import "testing"

func TestAdd(t *testing.T) {
	want := 10
	if got := Add(3, 7); got != want {
		t.Errorf("Hello() = %v, want %v", got, want)
	}
}
