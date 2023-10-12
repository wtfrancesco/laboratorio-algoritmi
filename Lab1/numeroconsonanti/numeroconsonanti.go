package main 

import (
		"fmt"
		"bufio"
		"os"
		"strings"

)
func numeroConsonanti(stringhe []string) {

	for _,stringa := range stringhe{
		count := 0 
		for _, runa := range stringa{
			if(strings.ContainsRune("bcdfghlmnpqrstvz", runa)){
				count++
			}
		}
		fmt.Println(stringa, "contiene ", count, "consonanti")
	}
}
func main (){

	scanner := bufio.NewScanner(os.Stdin)

	var stringhe []string

	for scanner.Scan(){
		stringhe = append(stringhe, scanner.Text())
	}

	numeroConsonanti(stringhe)


}