package main

import "fmt"

func main() {
	fmt.Println("Hello World")
}

// questions

// 1) how do we run the code in this project
// //////////////////////////////////////
// Answer: go run main.go
// /////////////////////////////////////

// 	//////////////// go cli commands ///////////////////////

// go build => compiles a bunch of go source code files
// go run => compiles and executes one or two files
// go fmt => formats all the code in each file in the current directory
// go install => compiles and installs a package
// go get => downloads the raw source code of someone else's package
// go test => runs any tests associated with the current project
// ///////////////////////////////////////////////////////

// 2) what does package main means
// //////////////////////////////////////
// Answer: package main is a special package that defines a standalone executable program instead of a library
// /////////////////////////////////////

// /////////////// 2 types of packages //////////////////////
// 1) executable => generates a file that we can run
// 2) reusable => code used as 'helpers'. good place to put reusable logic
// ///////////////////////////////////////////////////////

// 3) what does import "fmt" means
// /////////////////////////////////////////////////
// Answer: import "fmt" is a package in the Go standard library. It allows us to print things out in the terminal
// ////////////////////////////////////////////////

// 4) what does func main() means
// ///////////////////////////////////////////////
// Answer: func main() is the function that gets called when we run the executable program. It is the entry point for the executable program
// //////////////////////////////////////////////

// 5) how is the main.go file is organized
// ///////////////////////////////////////////////
// Answer: the main.go file is organized in a way that allows us to group together code that is related. In this case, we have one function, main, that is related to the other two lines of code. We can group them together with curly braces
// //////////////////////////////////////////////


