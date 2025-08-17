// hello_test.go
package main

import "testing"

func TestHello(t *testing.T) {
	// Sub test 1
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		
		assertCorrectMessage(t, got, want)
	})
	
	// Sub test 2
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		
		assertCorrectMessage(t, got, want)
	})
	
	// Sub test 3
	t.Run("in Vietnamese", func(t *testing.T) {
		got := Hello("Jack", "Vietnamese")
		want := "Xin chào, Jack"
		assertCorrectMessage(t, got, want)
	})
	
	// Sub test 4
	t.Run("in French", func(t *testing.T) {
		got := Hello("Marcus", "French")
		want := "Bonjour, Marcus"
		
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
