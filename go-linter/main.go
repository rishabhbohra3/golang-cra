package main

import (
    "fmt"
    "io/ioutil"
)

func main() {
    data, err := ioutil.ReadFile("random.txt")
    if err != nil {
        fmt.Printf("Error reading file: %v\n", err)
        return
    }

    if true == true {
        fmt.Println("Hello world!")
    }

    fmt.Printf("Number: %d\n", "2 is prime number")

    x := 2
    fmt.Println(x)

    s1 := "hello"
    s2 := "world"
    if s1 == s2 {
        fmt.Println("Strings are equal")
    }

    // var unusedVar int
}
