package main

import (
		"fmt"
		"os"
		"bufio"
)


func main () {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("CTRL+D per terminare ")


	for(scanner.Scan()){
		flag := true
		line := scanner.Text()

		for _ , runa := range line{
			if(runa == 'a' || runa == 'e' || runa == 'i' || runa == 'o' || runa == 'u'){
				fmt.Println(string(runa))
				flag = false
				break
			}
		}
		if(flag){
			fmt.Println("la parola non contiene vocali")
		}
	}
}