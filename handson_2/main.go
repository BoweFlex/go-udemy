package main

import "fmt"

func main() {
    // Print these values as both binary & hexadecimal
    printers := [7]int{0,1,2,3,4,5,100}
    for i := 0; i < len(printers); i++ {
        fmt.Printf("%v as binary: %b \n", printers[i], printers[i])
        fmt.Printf("%v as hexadecimal: %x \n", printers[i], printers[i])
    }
}
