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
	bubbleSort()
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
	var num = 7
	a := []int{0, 0, 0, 0, 0, 0, 0, 0,0}
	fmt.Print(a, "initial array: \n ")
	for i := 0; i < num; i++ { //adds 5 values to array
		a[i] = rand.Intn(7)
		fmt.Println(a)
	}
	fmt.Println("Starting sort")
	var temp = 0
	var temp1 = 0;
	for i := 0; i < len(a); i++ {
		if (a[i] >= a[i+1]) {
			swap(temp1, temp, a, i)
		} else{
			temp1, temp := a[i], a[i+1]
			fmt.Println(temp1,temp)
			a[i] = temp
			a[i+1] = temp1
			fmt.Println(a)
		}
		fmt.Println("val of ", i)
	}
}

func swap(temp1 int, temp int, a []int, i int) {
	temp1, temp = a[i+1], a[i]
	fmt.Println(temp1, temp)
	a[i] = temp1
	a[i+1] = temp
	fmt.Println(a)
}
