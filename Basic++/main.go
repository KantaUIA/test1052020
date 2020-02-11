package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//hWorld100()
	//enTilHundre(1, 300)
	//favFrukt()
	//whiletelling()

	//bubbleSort()
	betterSort()
	evenBetterSort()
}

//Oppgave 1 Lag en funksjon som printer “Hello” 100 ganger.  Hint: FOR

func hWorld100() {

	hello := "Hello world!"

	for i := 1; i < 100; i++ {
		fmt.Println(i, hello)
	}
}

//Oppgave 2 Lag en funksjon som teller fra et tall til et annet tall
func enTilHundre(i, j int) {

	for i := 1; i < j; i++ {
		fmt.Println(i)
	}
}

//Oppgave 3 & 4 Lag en array og fyll denne med navnet til dine 3 favorittfrukter

func favFrukt() {

	//kort declaration
	fav := [3] string{"Banan", "Mango", "Eple"}
	fmt.Println(fav)

	//lengere versjon
	var minFav [3] string
	minFav[0] = "Cheeseburger"
	minFav[1] = "Crispychicken"
	minFav[2] = "Angusburger"

	fmt.Println(minFav)

	//Oppgave 5 Utvid funksjonen og print EN tilfeldig frukt.
	//init ran gen
	rand.Seed(time.Now().UnixNano())
	velg := minFav[rand.Intn(len(minFav))]
	fmt.Println("Random burger = ", velg)

}

//Oppgave 6 whileloop telling?
func whiletelling() {

	i := 0

	for i < 100 {
		i++
		fmt.Println("Teller ", i)
	}
}

//Oppgave 7 Lag en array med 50 tall og sorter disse. Hint: sorteringspakke?

//func radomNumgen(){
//	for a := []; i := 0; i<50; i++{
//		a[i]=i
//	}
//
//	for (var a=[], i = 0; i<50; i++) a[i]=i;
//
//}
func bubbleSort() {
	var a [10]int
	fmt.Print(a, "initial array: \n ")
	for i := 0; i < len(a); i++ { //adds 5 values to array
		a[i] = rand.Intn(7)
		fmt.Println(a)
	}
	fmt.Println("Starting sort")
	var temp = 0
	var temp1 = 0;
	for i := 0; i < (len(a) * 2); i++ {
		if (i == len(a)) {
			i = 0
		}
		if (a[i] >= a[i+1]) {
			swap(temp1, temp, a, i)
		} else {
			temp1, temp := a[i], a[i+1]
			fmt.Println(temp1, temp)
			a[i] = temp
			a[i+1] = temp1
			fmt.Println(a)
		}
		fmt.Println("val of ", i)
	}
}

func swap(temp1 int, temp int, a [10]int, i int) {
	temp1, temp = a[i+1], a[i]
	fmt.Println(temp1, temp)
	a[i] = temp1
	a[i+1] = temp
	fmt.Println(a)
}
func betterSort() {
	list := createlist();
	fmt.Print(list, "initial array: \n ")
	fmt.Println("startin sort")
	len := len(list)
	for i := 0; i < len; i++ {
		for j := 0; j < len-1; j++ {
			if list[j] > list[j+1] {
				temp := list[j+1]
				list[j+1] = list[j]
				list[j] = temp
				fmt.Println(list)
			}
		}
	}
}

func createlist() []int {
	var list =[]int {0,0,0,0,0,0,0,0,0,0}
	for i := 0; i < len(list); i++ { //adds 5 values to array
		list[i] = rand.Intn(100)
		fmt.Println(list)

	}
	return list
}

func evenBetterSort() {
	list := createlist()
	fmt.Println(list)
	n := len(list)
	swapped := true
	for swapped {
		swapped = false
		for i := 1; i < n; i++ {
			if list[i-1] > list[i] {
				fmt.Println("swapping", list)
				list[i-1], list[i] = list[i], list[i-1]
				swapped = true
			}
		}
	}
	fmt.Println(list)
}
