package main

import "fmt"

type ByteSize int

const (
    first = iota // iota default == 0
    second
    third
    fourth
    fifth
    sixth
)

const (
    _ = iota
    KB ByteSize = 1 << (10 * iota)
    MB
    GB
    TB
    PB
    EB
)

func main() {
    /* // modify these to use iota
    fmt.Printf("%d \t %b \n", 1, 1)
    fmt.Printf("%d \t %b \n", 1<<1, 1<<1)
    fmt.Printf("%d \t %b \n", 1<<2, 1<<2)
    fmt.Printf("%d \t %b \n", 1<<3, 1<<3)
    fmt.Printf("%d \t %b \n", 1<<4, 1<<4)
    fmt.Printf("%d \t %b \n", 1<<5, 1<<5)
    fmt.Printf("%d \t %b \n", 1<<6, 1<<6)
    */

    fmt.Printf("%d \t %b \n", 1, 1)
    fmt.Printf("%d \t %b \n", 1<<first, 1<<first)
    fmt.Printf("%d \t %b \n", 1<<second, 1<<second)
    fmt.Printf("%d \t %b \n", 1<<third, 1<<third)
    fmt.Printf("%d \t %b \n", 1<<fourth, 1<<fourth)
    fmt.Printf("%d \t %b \n", 1<<fifth, 1<<fifth)
    fmt.Printf("%d \t %b \n", 1<<sixth, 1<<sixth)

    //Show the sizes of each measurement of bytes in DECIMAL and BINARY
    fmt.Printf("%d \t %b \n", KB, KB)
    fmt.Printf("%d \t %b \n", MB, MB)
    fmt.Printf("%d \t %b \n", GB, GB)
    fmt.Printf("%d \t %b \n", TB, TB)
    fmt.Printf("%d \t %b \n", PB, PB)
    fmt.Printf("%d \t %b \n", EB, EB)
}
