package main

import "fmt"

func main() {
	ch := make(chan string, 2)
	go ReadMessages(ch)
	go PrintMessages(ch)
	close(ch)
}

func ReadMessages(source chan<- string) {
	// writes messages into the channel
	source <- "New Message"
}

func PrintMessages(source <-chan string) {
	// prints messages
	fmt.Println(<-source)
}
