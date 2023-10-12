package main

import (
		"fmt"
		"bufio"
		"os"
)

func main () {

	scanner := bufio.NewScanner(os.Stdin)
	countA := 0
	countB := 0
	n := 5;
	for i := 0 ; i < n ; i++ {	
		if(scanner.Scan()){
			str := scanner.Text()
			if([]rune(str)[0] == 'a' || []rune(str)[0] == 'A'){
				countA++
			} else if([]rune(str)[0] == 'b' || []rune(str)[0] == 'B'){
				countB++
			}
		}
	}
	fmt.Println(" a: ", countA, "\n", "b: ", countB)
}