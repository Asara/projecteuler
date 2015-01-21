package main

import "fmt"

func main() {
    start := 2
    evenly_divisible := 0

    for (evenly_divisible == 0) {
       if (divisible(start)){ evenly_divisible = start } 
    }
    fmt.Println(evenly_divisible)
}

func divisible(n int) bool {
}
