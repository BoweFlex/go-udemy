// This will fail, because channels block routines by default.
// Values are required to be pulled back off before closing the initial connection.
package main

import "fmt"

func main() {
    c := make(chan int)

    c <- 42 // causes "all goroutines are asleep error"

    fmt.Println(<-c)
}
