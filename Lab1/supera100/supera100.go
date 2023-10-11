package main

import(
	"fmt"
)

func main () {
	
	input := 0; 
	supera100 := 0; 
	flag := true; //mi serve per inserire il valore in supear100 solamente una volta

	for {

		fmt.Scan(&input);
		
		if(input == -1){
			break;
		}
		
		if(input > 100 && flag == true){
			supera100 = input;
			flag = false;
		}		

	}

	if(supera100 != 0){
		fmt.Println(supera100)		
	}else{
		fmt.Println("nessun numero maggiore di 100")
	}
	

}
