package main

import "fmt"

func signalChanger(num int) int {
	return num * -1
}

func signalChangerWithPointer(num *int) {
	*num = *num * -1
}

func main() {
	num := 20
	oposite_signal_num := signalChanger(num)
	fmt.Println(oposite_signal_num)
	fmt.Println(num)

	num2 := 40
	signalChangerWithPointer(&num2)
	fmt.Println(num2)

}