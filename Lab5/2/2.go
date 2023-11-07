//Come deve essere completato il caso base? 
//largest(numbers[:n-1])

//Durante l’esecuzione della chiamata largest(numbers), considerate la chiamata ricorsiva che termina per prima. Qual è l’argomento passato in questa chiamata e quale valore restituisce la funzione?
//L'argomento passato alla chiamata ricorsiva che termina per prima è la slice numbers che contiene solamente un elemento [1]

//Con quali argomenti viene eseguita per la prima volta la funzione max? 
//Viene eseguita per la prima volta con 2 ed 1

//Con quali argomenti viene eseguita l'ultima volta la funzione? 
//L'ultima volta la funzione viene eseguita con 8 e 21  

package main

import (
	"fmt"
)

func max(x, y int) int {
	if x > y {
		return x

	} else {
		return y
	}

}

func largest (numbers []int) int{
	
	n := len(numbers)

	if n == 1  {
		return numbers[0]
	}

	return max(numbers[n-1], largest(numbers[:n-1]))
}

func main() {

	numbers := []int{1,2,5,7,-2,10,9,21,3,8}

	n := largest(numbers)
	fmt.Println(n)


}

