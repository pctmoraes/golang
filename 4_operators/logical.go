package main

import "fmt"

func main() {
    var x bool = true
    var y bool = false
    var z bool = false

    fmt.Println(x && y)
    fmt.Println(!x && !y)

    fmt.Println(x || y)
	fmt.Println(!x || !y)
	fmt.Println(!x || !y && z)
}
