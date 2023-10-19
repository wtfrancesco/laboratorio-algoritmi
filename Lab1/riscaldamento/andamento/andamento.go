package main

import (
	"fmt"
)

func main () {
	
	prev := 0
	next := 0

	fmt.Scan(&prev)

	for{
		
		fmt.Scan(&next)

		if(next == 0){
			break
		}
		
		if(next >= prev){
			fmt.Println("+")
		}else{
			fmt.Println("-")
		}

		prev = next 

	}	
	

}
