package main

import "fmt"

func main() {
    fmt.Println(Hello("Chris"))
}

//func Hello() string {
//    return "Hello, world!"
//}
//

const englishHelloPrefix = "Hello, "

func Hello(name string) string {
    return fmt.Sprintf("%v%v", englishHelloPrefix, name)
}
