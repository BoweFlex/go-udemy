// Can add information to errors with:
// builtin.error or errors.New() (or fmt.Errorf())
/* “Error values in Go aren’t special, they are just values like any other, 
and so you have the entire language at your disposal.” - Rob Pike */
// Example 1: Return error along with desired return from func.

package main

import (
	"errors"
	"log"
)

var ErrNorgateMath = errors.New("norgate math: square root of negative number")

func main() {
	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("norgate math: square root of negative number")
		// can return the same thing but with a variable
		return 0, ErrNorgateMath
	}
	return 42, nil
}
