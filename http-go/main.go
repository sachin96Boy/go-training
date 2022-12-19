package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// create a custom type associate with a function
type logWriter struct{}

func main() {
	fmt.Println("Hello World")
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// create a new instance of logWriter
	lw := logWriter{}
	io.Copy(lw, resp.Body)

	// //////////////////////////////////////////////////
	// // copy the response body to os.Stdout
	// // io.Copy(os.Stdout, resp.Body)
	// //////////////////////////////////////////////////

	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	os.Exit(1)
	// }
	// fmt.Printf("%s", body)

	// fmt.Println(resp)
}

// create a method for logW riter
func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}
