package main

import (
	"fmt"
	"github.com/kmehran1106/learn-go-with-tests/hello_world"
)

func main() {
	fmt.Println(hello_world.Hello("Chris", "english"))
	fmt.Println(hello_world.Hello("Chris", "spanish"))
	fmt.Println(hello_world.Hello("Chris", "french"))
	fmt.Println(hello_world.Hello("", ""))
}
