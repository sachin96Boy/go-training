package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://golang.org",
		"http://stackoverflow.com",
		"http://amazon.com",
	}

	c := make(chan string) // create a channel

	for _, link := range links {
		go checkLink(link, c) // run the checkLink function as a go routine
	}

	// for {
	// 	go checkLink(<-c, c) // run the checkLink function as a go routine
	// this creates a infinite loop
	// }

	// for l := range c {
	// time.Sleep(5 * time.Second)
	// 	go checkLink(l, c) // run the checkLink function as a go routine
	// }

	// here when we add sleep to the for loop, the program will wait for 5 seconds before checking the next link
	// but this is a blocking operation, so the program will not check the other links until the 5 seconds are up  

	// fmt.Println(<-c) // receive the link from the channel

	for l := range c {
		go func(link string) { // anonymous function
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}

}

// checklink checks a link and prints if it is up or down

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link // send the link to the channel
		return
	}

	fmt.Println(link, "is up!")
	c <- link // send the link to the channel
}

// Go Routine
// A go routine is a lightweight thread of execution.
// A go routine is a function or method that is capable of running concurrently with other functions or methods.
// it runs in the background and does not block the execution of other functions or methods.

// channels in Go
// A channel is a conduit that allows two goroutines to communicate with each other and synchronize their execution.
// channels can be used to pass values of a specified type between goroutines.
// A channel is created using the built-in make function.
// The make function takes two arguments: the type of the channel and the buffer size.
// The buffer size is optional and is used to specify the size of the buffer.
// A buffer is a temporary storage area where data is kept while it is being transferred from one place to another.
// A buffer is used to increase the performance of an application by allowing it to read and write data in chunks.
// A channel can be buffered or unbuffered.

// sending data to a channel
// A channel can be used to send data to a go routine.
// The syntax for sending data to a channel is:
// channel <- value
// The value is sent to the channel and the program execution continues.
// The value is received from the channel using the syntax:
// <- channel
// The program execution is blocked until a value is received from the channel.
