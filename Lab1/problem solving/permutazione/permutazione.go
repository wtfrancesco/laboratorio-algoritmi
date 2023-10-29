package main

import (
		"fmt"
)

type person struct {
	id int
    name string
}

func main (){

	utenti := []person{person{id: 1, name: "Andrea"},person{id: 6, name: "Francesco"},person{id: 5, name: "Elisa"},person{id: 2, name: "Beatrice"},person{id: 3, name: "Carlo"},person{id: 4, name: "Dino"}, person{id: 7, name: "Giorgia"},person{id: 9, name: "Irene"},person{id: 8, name: "Henry"}}

	fmt.Println("Utenti prima: ", utenti, "\n")
	
	for i := 0 ; i < len(utenti) ; i++{
		for j := 0; j < len(utenti)-i-1 ; j++{
			if(utenti[j].id < utenti[j+1].id){
				utenti[j] , utenti[j+1] = utenti[j+1],utenti[j]
			}
		}
	}

	fmt.Println("Utenti dopo: ", utenti,"\n")
}