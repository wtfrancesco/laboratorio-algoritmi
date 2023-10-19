package main

import (
	"fmt"
)


func main () {
	var parola1, parola2 string 
	fmt.Print("Parola 1: ")
	fmt.Scan(&parola1)

	fmt.Print("Parola 2: ")
	fmt.Scan(&parola2)

	parola1Mappa := make(map[rune]int)
  	parola2Mappa := make(map[rune]int)

  	for _ , r := range parola1 {
  		parola1Mappa[r]++
  	}

  	for _ , r := range parola2 {
  		parola2Mappa[r]++
  	}

  	if(len(parola1Mappa) != len(parola2Mappa)){
  		fmt.Println("non sono anagrammi")
  	}else{
  		for k , _ := range parola1Mappa{
  			if(parola1Mappa[k] != parola2Mappa[k]){
  				return 
  			}
  		}
  		fmt.Println("sono anagrammi")
  	}


}