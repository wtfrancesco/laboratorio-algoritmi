package main

import (
	"fmt"
)

/*
4 2 1 

*/


func main (){
	var consumiEnergetici []int
	consumo := 0
	fmt.Println("9999 per terminare")
	for {
		fmt.Scan(&consumo)
		if(consumo == 9999){
			break
		}
		consumiEnergetici = append(consumiEnergetici, consumo)
	}

	for i := 1; i < len(consumiEnergetici)-1 ; i++{
		if(consumiEnergetici[i-1] > consumiEnergetici[i] && consumiEnergetici[i]<consumiEnergetici[i+1]){
			fmt.Print(i+1, " ")
		}
	}

	fmt.Println()

}