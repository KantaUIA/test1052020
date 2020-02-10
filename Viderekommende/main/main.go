package main

import (
	arr "10022020/test1052020/Viderekommende/getsliced"
	get "10022020/test1052020/Viderekommende/methodmethod"
	"fmt"
	"os"
)

func main() {

	args := os.Args

	if len(args) > 1 {
		fmt.Println(get.GetMessage(args[1]))
	} else {
		fmt.Println("Not enough args")
	}

	nums := arr.GetArrayStartingAt(5)
	numsHalf := nums[:5]
	fmt.Println(nums)
	fmt.Println(numsHalf)

}
