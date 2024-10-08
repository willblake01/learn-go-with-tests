package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Will")

	got := buffer.String()
	want := "Hello, Will"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
