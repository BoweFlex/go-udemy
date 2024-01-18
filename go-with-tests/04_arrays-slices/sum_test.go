package main

import (
    "testing"
    "slices"
   // "fmt"
)

func assertCorrectSum(t testing.TB, got, want int) {
    t.Helper()
    if got != want {
        t.Errorf("expected %d but got %d", want, got)
    }
}

func assertCorrectSumAll(t testing.TB, got, want []int) {
    t.Helper()
    if !slices.Equal(got, want) {
        t.Errorf("expected %v but got %v", want, got)
    }
}

func TestSum(t *testing.T) {
    t.Run("Sum of any count of numbers", func(t *testing.T) {
        numbers := []int{1,2,3}
        got := Sum(numbers)
        want := 6
        assertCorrectSum(t, got, want)
    })
}

func TestSumAll(t *testing.T) {
    t.Run("Sum of any number of slices", func(t *testing.T) {
        got := SumAll([]int{1,2}, []int{0,9})
        want := []int{3, 9}

        assertCorrectSumAll(t, got, want)
    })
}
func TestSumAllTails(t *testing.T) {
    // Locks scope to inside this function, and ensures type safety (can't accidentally pass a string and get a runtime error)
    assertCorrectSumAllTails := func(t testing.TB, got, want []int) {
        t.Helper()
        if !slices.Equal(got, want) {
            t.Errorf("expected %v but got %v", want, got)
        }
    }
    t.Run("Sum of any number of slices", func(t *testing.T) {
        got := SumAllTails([]int{1,2}, []int{0,9})
        want := []int{2, 9}

        assertCorrectSumAllTails(t, got, want)
    })
    t.Run("Safely handle empty slices", func(t *testing.T) {
        got := SumAllTails([]int{}, []int{0,9})
        want := []int{0, 9}

        assertCorrectSumAllTails(t, got, want)
    })
}
