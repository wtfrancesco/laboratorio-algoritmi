package main

import (
	"fmt"
)


func main (){

	current := 0
	prev := 0
	coppieUguali := 0

	fmt.Scan(&prev)

	for{
		fmt.Scan(&current)

		if(current == 0){
			break
		}

		if(prev == current){
			coppieUguali++
		}
		prev = current
	}

	fmt.Println(coppieUguali)



}