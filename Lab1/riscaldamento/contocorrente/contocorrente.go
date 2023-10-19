package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1{
		fmt.Println("Manca il saldo iniziale")
		return 
	}
	
	saldo , _ := strconv.Atoi(os.Args[1])
	
	var spesa int 

	for saldo > 0{

		fmt.Scan(&spesa)
		saldo -= spesa
		
	
	}
	
	
}
