package HelloWorld

import "fmt"

// print hello, world func 
func Hello(name string) string{
    return "Hello, " + name
}

// print hello, world
func main() {
    fmt.Println(Hello("World"))
}