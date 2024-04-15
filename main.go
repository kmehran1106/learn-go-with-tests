package main

import (
	"fmt"
	"github.com/kmehran1106/learn-go-with-tests/hello_world"
	"github.com/kmehran1106/learn-go-with-tests/integers"
)

func main() {
	// for the sake of testing we are running hello world chapter
	fmt.Println(hello_world.Hello("Chris", "english"))
	fmt.Println(hello_world.Hello("Chris", "spanish"))
	fmt.Println(hello_world.Hello("Chris", "french"))
	fmt.Println(hello_world.Hello("", ""))

	// for the sake of testing we are running integer chapters here
	fmt.Println(integers.Add(2, 2))
}
