package main 

import (
		"fmt"
)

func sassi(height int) int {
    if height == 1 {
        return 1
    }

    lowerSquares := height * height

    previousPyramid := sassi(height - 1)

    return lowerSquares + previousPyramid
}

func main(){
	var n int 
	fmt.Scan(&n)
	s := sassi(n)
	fmt.Println(s)
}