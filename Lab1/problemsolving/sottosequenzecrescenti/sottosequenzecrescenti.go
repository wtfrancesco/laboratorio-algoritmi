package main

import (
	"fmt"
	"errors"
	"io"
)

func main (){

	prev := 0 
	current := 0
	fmt.Scan(&prev)
	flag := false
	minimi := 0


	for{

		_ , err := fmt.Scan(&current)

		if err != nil {
        	if errors.Is(err, io.EOF) { // prefered way by GoLang doc
            	fmt.Println("Reading file finished...")
        	}
        	break
    	}

		if(prev < current){
			flag = true
		}else{
			if flag{
				minimi++
				flag = false
			}
		}
		prev = current
	}

	fmt.Println(minimi)
}
