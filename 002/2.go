package main

import "fmt"


func main() {
    x   := 1
    y   := 1
    z   := 0
    sum := 0
    for z < 4000000 {
        z = x + y
        if (z % 2 == 0) {
            sum = sum + z
        }
        x = y
        y = z
    }
    fmt.Println(sum)
}
