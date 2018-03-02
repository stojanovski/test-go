package main

import "fmt"

func main() {
    m := make(map[string]int)

    m["Answer"] = 42
    fmt.Println("The value:", m["Answer"])

    m["Answer"] = 48
    fmt.Println("The value:", m["Answer"])

    delete(m, "Answer")
    fmt.Println("The value:", m["Answer"])

    v, ok := m["Answer"]
    fmt.Println("The value:", v, "Present?", ok)

    n := make(map[string]int)
    i, j := n["one"]
    fmt.Println(n, i, j)
    n["one"] = 1
    i, j = n["one"]
    fmt.Println(n, i, j)

    o := make(map[int]string)
    //o[0] = "zero"
    if o == nil {
        fmt.Println("nil")
    }
    fmt.Println(o[0], "AA")
    fmt.Println(m["fifty"])
}

