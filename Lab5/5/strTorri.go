package main

import (
    "fmt"
)

func main () {
  var n int
  fmt.Scan(&n)
  torri := []string{"","",""}
  for i:= 65 ; i < 65+n ; i++{
    torri[0] += string(i)
  }
  fmt.Println(torri[0],",",torri[1],",",torri[2])
  hanoi(n,0,1,2,torri)
}

func hanoi (n, from, temp, to int, torri []string) {
  if n > 0 {
    hanoi(n-1,from,to,temp,torri) //sposto 2 dischi al temporaneo (immagino come se si spostassero in modo giusto)
    tmp:=torri[from][len(torri[from])-1]
    torri[from] = torri[from][:len(torri[from])-1]
    torri[to]=torri[to]+string(tmp)
    fmt.Println(torri[0],",",torri[1],",",torri[2],n)
    hanoi(n-1,temp,from,to,torri) //sposto dal temporaneo a quello finale
  }
}

