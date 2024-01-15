package main

import "testing"

//Old test
//func testHello(t *testing.T) {
//    got := Hello()
//    want := "Hello, world!"
//
//    if got != want {
//        t.Errorf("got %q want %q", got, want)
//    }
//}
func TestHello(t *testing.T) {
	got := Hello("Chris")
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
