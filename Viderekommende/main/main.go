package main

import(
	"os"
	"fmt"
	get "Viderekommende/methodmethod"
)

func main() {

	args := os.Args

	if len(args) > 1 {
		fmt.Println(get.GetMessage(args[1]))
	} else {
		fmt.Println("Not enough args")
	}

}