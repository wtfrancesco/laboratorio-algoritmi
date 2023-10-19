package main

import (
		"fmt"
		"bufio"
		"os"
		"unicode"
		"sort"
)

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	line := ""
	if(scanner.Scan()){
		line = scanner.Text()
	}

	fmt.Println(line)
	var m map[string]int
	m = make(map[string]int)
	
	for _ , runa := range line{
		if(runa >= 'A' &&  runa <= 'Z' || runa >= 'a' && runa <= 'z' ){
			m[string(unicode.ToUpper(runa))]++
		}
	}

	var keys []string

	for k := range m{
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _,value := range keys {
		fmt.Print(value, " ")
		for i := 0 ; i < m[value] ; i++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

}