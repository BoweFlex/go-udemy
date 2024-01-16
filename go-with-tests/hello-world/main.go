package main

import "fmt"

func main() {
}

//func Hello() string {
//    return "Hello, world!"
//}
//

func Hello(name string, language string) string {
    languages := map[string]string {
        "English": "Hello, ",
        "Spanish": "Hola, ",
        "French": "Bonjour, ",
    }
    var greeting string
    if name == "" {
        name = "World"
    }
    for k, v := range languages {
        if k == language {
            greeting = v
        }
    }
    if greeting == "" {
        greeting = languages["English"]
    }
    return fmt.Sprintf("%v%v", greeting, name)
}
