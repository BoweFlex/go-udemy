// Method set practice
package main

import (
    "fmt"
)

type person struct {
    name string
    age int
}

func (p *person) speak() {
    fmt.Printf("Hello, my name is %v and I am %v years old.", p.name, p.age)
}

type human interface {
    speak()
}

func saySomething(h human) {
    h.speak()
}

func main() {
    your_mom := person{
        "Mother",
        100,
    }
    your_mom_points := &your_mom

    //saySomething(your_mom) // this should fail
    saySomething(your_mom_points) // this should work
    
}
