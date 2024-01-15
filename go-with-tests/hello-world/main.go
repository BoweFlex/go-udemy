package main

import "fmt"

func main() {
    fmt.Println(Hello("Chris"))
}

//func Hello() string {
//    return "Hello, world!"
//}
//
func Hello(name string) string {
    return fmt.Sprintf("Hello, %v", name)
}
