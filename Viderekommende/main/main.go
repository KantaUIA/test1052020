package main

import (
	board "10022020/test1052020/Viderekommende/board"
	arr "10022020/test1052020/Viderekommende/getsliced"
	get "10022020/test1052020/Viderekommende/methodmethod"
	"fmt"
	"os"
)

func main() {
	arraydemo(5)
	b := board.CreateBoard()
	fmt.Println(b[0])
	fmt.Println(b[1])
	fmt.Println(b[2])
}

func greetings() {
	args := os.Args

	if len(args) > 1 {
		fmt.Println(get.GetMessage(args[1]))
	} else {
		fmt.Println("Not enough args")
	}
}

func arraydemo(i int) {
	nums := arr.GetArrayStartingAt(i)
	numsHalf := nums[:len(nums)/2]
	fmt.Println(nums)
	fmt.Println(numsHalf)
}
