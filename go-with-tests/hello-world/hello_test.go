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

//One test grouped around a "thing" and then defined subtests for different scenarios
func TestHello(t *testing.T) {
    t.Run("saying hello to people", func(t *testing.T) {
        got := Hello("Chris")
        want := "Hello, Chris"

        if got != want {
            t.Errorf("got %q want %q", got, want)
        }
    })
    t.Run("say 'Hello, World' when empty string is supplied", func(t *testing.T) {
        got := Hello("")
        want := "Hello, World"

        if got != want {
            t.Errorf("got %q want %q", got, want)
        }
    })
}

