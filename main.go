package main

import (
    "fmt"
    "io/ioutil"
)

func main() {
    ioutil.ReadFile("random.txt")

    if true == true {
        fmt.Println("Hello world!")
    }

    fmt.Printf("Number: %d\n", "2 is prime number")

    x := 0
    x = 1 
    x = 2
    fmt.Println(x)

    s1 := "hello"
    s2 := "world"
    if fmt.Sprintf("%s", s1) == fmt.Sprintf("%s", s2) {
        fmt.Println("Strings are equal")
    }

    // var unusedVar int
}
