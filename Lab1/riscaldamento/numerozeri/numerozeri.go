package main

import (
		"fmt"
)

func main(){
	var valori []int 
	input := 0;
	for {
		fmt.Scan(&input)
		if (input == -1){
			break
		}
		valori = append(valori, input)
	}

	indiceMax := 0
	maxZeri := 0
	for i , numeri := range valori{
		countZeri := 0
		for numeri != 0{
			if(numeri % 10 == 0){
				countZeri++
			}
			numeri = numeri / 10 
		}
		if maxZeri < countZeri{
			maxZeri = countZeri
			indiceMax = i
		}
	}
	fmt.Println(valori[indiceMax])
}