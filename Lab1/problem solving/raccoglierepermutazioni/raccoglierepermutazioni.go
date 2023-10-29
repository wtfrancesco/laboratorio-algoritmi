package main

import "fmt"

func main() {
    //permutation := []int{4, 5, 1, 3, 6, 2}
    permutation := []int{6, 5, 4, 3, 2, 1}

    count := 0
    max := 0

    for i := 0; i < len(permutation); i++ {
       
    	if(permutation[i] > max){
    		max = permutation[i]
    	}else{
    	 count++
    	 max = permutation[i]
    	}


    }

    fmt.Printf("Il numero di volte in cui bisogna tornare verso sinistra è: %d\n", count)
}
