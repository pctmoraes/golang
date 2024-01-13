package main

import "fmt"

func main() {
    var x = 10
    
	x++
    fmt.Println(x)

	x += 23
    fmt.Println(x)

	x--
    fmt.Println(x)

	x -= 19
    fmt.Println(x)

	x *= 2
    fmt.Println(x)
}
