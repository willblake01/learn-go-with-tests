package helloworld

import (
	"fmt"
	"testing"
)

func assertCorrectMessage(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestHello(t *testing.T) {
	t.Run("saying hello to people in English", func(t *testing.T) {
		got := Hello("Will", "English", "!")
		want := "Hello, Will!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish", "!")
		want := "Hola, Elodie!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Elodie", "French", "!")
		want := "Bonjour, Elodie!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, world!' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "", "!")
		want := "Hello, world!"

		assertCorrectMessage(t, got, want)
	})
}

func ExampleHello() {
	sum := Hello("Will", "English", "!")
	fmt.Println(sum)
	// Output: Hello, Will!
}
