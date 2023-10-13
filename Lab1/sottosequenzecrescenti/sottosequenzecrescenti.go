//diEOF.Adesempio,suinput1 2 13 0 7 8 9 -1 0 2ilprogrammastampalesomme16,24 e 1.
//non mi serve salvare 
package main

import (
		"fmt"
)

func main(){
	current := 0;
	prev:= 0
	fmt.Scan(&prev)
	somme := prev
	
	for{
		_, err := fmt.Scan(&current)
	
		if(err != nil){
			if(somme != 0){
				fmt.Println(somme)
			}
			break
		}

		if(current > prev){
			somme += current
		}else{
			fmt.Println(somme)
			somme = current
		}
		
		prev = current	
	}
}