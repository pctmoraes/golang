package main

import "fmt"

func main() {
	channel := make(chan string, 200)
	channel <- "Olá Mundo!"
	channel <- "Programando em Go!"
	channel <- "Programando em Go De Novo!"

	message := <-channel
	message2 := <-channel

	fmt.Println(message)
	fmt.Println(message2)
}

// Buffered Channels:

// A buffered channel, on the other hand, has a specified capacity, allowing it to hold a certain number of elements without blocking the sender. The sender can send up to the buffer capacity without waiting for the receiver.
// When the buffer is full, further send operations will block until there is space in the buffer.
// The receiver can retrieve elements from the channel without waiting for the sender as long as there are items in the buffer. Once the buffer is empty, further receive operations will block until the sender adds more items.
// Buffered channels are useful when you want to decouple the sender and receiver to some extent, allowing them to operate at different rates.