package main

import (
	_"reflect"
	"fmt"
)

func quanteConA(parole []string) int{
	numeroA := 0;
	for _ , str := range parole{
		for _ , char := range str{
			if(char == 'a'){
				numeroA++
				break
			}
		}
	}
	return numeroA
}

func containsLetter(char rune, str string) bool{
	for _, actualChar := range str{
		if(actualChar == char){
			return true
		}
	}

	return false
}

func primaConA(parole []string) string{
	for _,v := range parole{
		if containsLetter('a',v){
			return v
		}
	}
	return ""
}

func piuCorta(parole []string) int{
	min := len(parole[0])

	for i := 1 ; i < len(parole) ; i++ {
		if len(parole[i])<min{
			min = len(parole[i])
		}
	}
	return min
}


func main() {
	const N = 3
	parole := make([]string, N)
	
	for i := 0; i < N; i++ { 
		fmt.Scan(&parole[i])
	}

	fmt.Println(quanteConA(parole)) 
	fmt.Println(primaConA(parole)) 
	fmt.Println(piuCorta(parole)) 
}
