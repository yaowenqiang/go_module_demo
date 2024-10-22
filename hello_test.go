package hello

import "testing"

func TestHello(t *testing.T) {
	want := "hello, world."
	if got := hello(); got != want {
		t.Errorf("hello() = %q, want %q", got, "hello")
	}
}
