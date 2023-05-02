package HelloWorld

import "fmt"

const englishHelloPrefix = "Hello, "

// print hello, world func 
func Hello(name string) string{
    if name == ""{
        return englishHelloPrefix + "World"
    }
    
    return englishHelloPrefix + name
}

// print hello, world
func main() {
    fmt.Println(Hello("World"))
}