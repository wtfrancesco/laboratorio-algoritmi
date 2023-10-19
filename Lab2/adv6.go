/*
Input Prof:
18 Giorni 1638
80 Giorni 362740
256 Giorni 1644874076764
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	tempPesci := []string{}
	for scanner.Scan() {
		tempPesci = strings.Split(scanner.Text(), ",")
	}

	mappaPesci := make(map[int]int)

	for i := 0; i < len(tempPesci); i++ {
		val, _ := strconv.Atoi(tempPesci[i])
		mappaPesci[val]++
	}

	for tempo := 0; tempo < 256; tempo++ {
		temp0 := mappaPesci[0]
		for i := 0; i < 8; i++ {
			mappaPesci[i], mappaPesci[i+1] = mappaPesci[i+1], mappaPesci[i]
		}
		mappaPesci[6] = mappaPesci[6] + temp0
	}
	somme := 0
	for _, v := range mappaPesci {
		somme += v
	}

	fmt.Println(somme)

}
