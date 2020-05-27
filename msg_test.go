package minipkg

import "testing"

func TestHello(t *testing.T) {
	want := "hello"
	if got := HelloMessage(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}

func TestRateLimit(t *testing.T) {
	MockFireAndRecv()
}
