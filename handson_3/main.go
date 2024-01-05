package main

import "fmt"

func main() {
    y := 42
    z := 42.0 // Assumes float64
    var m float32 = 43.742
    // Check your types
    fmt.Printf("%v of type %T \n", y, y)
    fmt.Printf("%v of type %T \n", z, z)
    fmt.Printf("%v of type %T \n", m, m)


    // You can't store a float32 and store it in a float64 var, or int, etc.
    // z = m
    fmt.Printf("%v of type %T \n", z, z)

    // Casting values
    // explicit conversion
    z = float64(m)
    fmt.Printf("%v of type %T \n", z, z)

}
