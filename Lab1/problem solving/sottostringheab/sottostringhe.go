package main


import (
		"fmt"
)

func main(){
	stringa := "ccbaacbabbcbab"
	//stringa := "ababab"
	//stringatest := "x"
	//fmt.Println(stringatest[0:1])
	sottostringhe := 0
	for i,runa := range stringa {
		if runa == 'a'{
			for x, runa2 := range stringa[i:]{
				if(runa2 == 'b'){
					fmt.Println(stringa[i:i+x+1])
					sottostringhe++
				}
			}
		}
	}
	fmt.Println(sottostringhe)
}