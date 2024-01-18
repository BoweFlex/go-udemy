package iteration

import (
    "testing"
    "fmt"
)

func assertCorrectRepeat(t testing.TB, got, want string) {
    t.Helper()
    if got != want {
        t.Errorf("expected %q but got %q", want, got)
    }
}

func TestRepeat(t *testing.T) {
    t.Run("Repeating a", func(t *testing.T) {
        repeated := Repeat("a", 5)
        expected := "aaaaa"
        assertCorrectRepeat(t, repeated, expected)
    })
    t.Run("Repeating 'a' ten times", func(t *testing.T) {
        repeated := Repeat("a", 10)
        expected := "aaaaaaaaaa"
        assertCorrectRepeat(t, repeated, expected)
    })
}

func ExampleRepeat() {
    repeated := Repeat("jbowe rocks ", 2)
    fmt.Println(repeated)
    // Output: jbowe rocks jbowe rocks 
}

func BenchmarkRepeat(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Repeat("a", 5)
    }
}
