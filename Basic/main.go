package main

import (
	"fmt"
	"math/rand"
	
)

func main() {
	hWorld()
	sum()
	returnArgs("word", "word2")
	ranNum()
	ranAlder()
	sjekkTallet(20)
	sjekkString("hei")
	sjekkString("ok")

}

//Oppgave 1 
func hWorld(){
	a := "Hello "
	b := "Kanta!"
	fmt.Println(a + b)
}


//Oppgave 2 Lag en FUNKSJON som returnerer summen av to tall. Summen skal printes i func main.

func sum() {
	a := 100
	b := 200
	fmt.Println("Sum =", a+b)
}

//Oppgave 3 Lag en funksjon som tar imot to argumenter fra kommandolinjen og printer disse. Hint: Package os
func returnArgs(a, b string ){
	fmt.Println(a + " " + b)
}

//Opppgave 4 Lag en funksjon som generer et random tall og printer dette

func ranNum() {
	fmt.Println("Random nr =", rand.Intn(100))
}


//Oppgave 5 Utvid funksjonen til å generere et random tall mellom 0 og alderen din. 
func ranAlder() {
	fmt.Println("Random mellom alderen =", rand.Intn(22))
}

//Oppgave 6 Lag en funksjon som sjekker om tallet er høyere eller lavere enn 15. Om det er høyere: Print “ Over 15!” Om det er lavere: “ Under 15”!
func sjekkTallet(a int) {

	if a >= 15 {
		fmt.Println("Større enn 15!")
	} else {
		fmt.Println("Mindre enn 15!")
	}
}

// sier hei tilbake om du sier hei, eller sier den nei!

func sjekkString(a string){
	
	b := "hei"

	if (a == b) {
		fmt.Println("Hei på deg!")
	}else {
		fmt.Println("Nei!")
	}
}


