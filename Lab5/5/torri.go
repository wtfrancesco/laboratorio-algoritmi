package main

import (
    "fmt"
)

func main () {
  fmt.Print("Inserisci n > 0: ")
  var n int
  fmt.Scan(&n)
  hanoi(n,0,1,2) //from, temp, to
  nb :=  Nhanoi(n,0,1,2)
  fmt.Println(nb)
}

func hanoi (n, from, temp, to int) {
  if n == 1 {
    fmt.Println("Muovo", n, "da: ", from, "->", to)
    return
  }
    hanoi(n-1,from,to,temp) //sposto 2 dischi al temporaneo (immagino come se si spostassero in modo giusto)
    fmt.Println("Muovo",n ,"da: ", from ,"->", to)
    hanoi(n-1,temp,from,to) //sposto dal temporaneo a quello finale
}


func Nhanoi (n, from, temp, to int) int{
    cont  := 0
  if n > 0 {
    cont++
    cont += Nhanoi(n-1,from,to,temp)
    cont += Nhanoi(n-1,temp,from,to) //1,0,2
  }
  return cont
}

