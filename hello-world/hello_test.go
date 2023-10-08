package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("", func(t *testing.T) {
		got := Hello("Niklas")
		want := "Hello, Niklas"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
	t.Run("", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
