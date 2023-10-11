/*TESTO ESERCIZIO
Scrivete un programma che legga una riga di caratteri terminata da a-capo, poi legga un’altra lettera e stampi quante volte la lettera compare nella prima riga.
*/

package main

import (
  "fmt"
  "bufio"
  "os"
)

func main(){
    scanner := bufio.NewScanner(os.Stdin)
    line := ""
    occorrenze := 0
    if(scanner.Scan()){
      line = scanner.Text()
    }else{
      fmt.Println("errore in lettura")
    }

    tempLetteraDaContare := "";
    var letteraDaContenere rune

    fmt.Scan(&tempLetteraDaContare)

    letteraDaContenere = []rune(tempLetteraDaContare)[0]

    for _ , runa := range line{
      if(runa == letteraDaContenere){
        occorrenze++
      }
    }

    fmt.Println(occorrenze)

}
