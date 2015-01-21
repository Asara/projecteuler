package main

import (
	"fmt"
)

func main() {
	search := 600851475143
	f := 2
	for f*f <= search {
		if search%f == 0 {
	/*		fmt.Printf("%s\n", f)*/
			search = search / f
	} else {
			f = f + 1
		}
	}
    fmt.Printf("%s\n", search)
}
