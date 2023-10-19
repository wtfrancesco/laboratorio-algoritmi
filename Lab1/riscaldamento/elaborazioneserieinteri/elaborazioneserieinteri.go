package main

import (
	"fmt"
)

func stranoProdotto(numeri []int) int{
	prod := 1 
	for i := 0 ; i < len(numeri) ; i++{
		if(numeri[i] % 4 == 0 && numeri[i] > 7){
			prod *= numeri[i]	
		}	
	}


	/*Soluzione con for range 
		//ho scritto key e value anche se non li tutti e due ma per ricordarmi
		//come era il for range su le slice in go
		for key , value := range numeri{
			
		}

	*/
	
	return prod
}


func pariDispari(numeri []int){
	for _ , n := range numeri{
		if n % 2 == 0{
			fmt.Println(n, "è pari")
		}else{
			fmt.Println(n, "è dispari")
		}
	}
}


func minDispari(numeri []int) int{
	min := 0;
	i := 0  
	for i = 0 ; i < len(numeri) ; i++{
		if numeri[i] % 2 != 0{
			min = numeri[i]
			break
		}
	}

	for j := i+1 ; j < len(numeri) ; j++{
		if numeri[j] < min{
			min = numeri[i]
		}
	}

	return min
}


func main () {

	const N = 10
	numeri := make([]int, N)
	
	for i := 0 ; i < N ; i++{
	
		fmt.Scan(&numeri[i])
	
	}
	
	fmt.Println(stranoProdotto(numeri))
	pariDispari(numeri)
	fmt.Println(minDispari(numeri))

}
