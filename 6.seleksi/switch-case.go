package main 

import "fmt"

func main(){

	pilih := 1

	switch pilih {
		case 1 : 
			fmt.Println("ini pilihan ke-1")
		case 2 : 
			fmt.Println("ini pilihan ke-2")
		case 3 : 
			fmt.Println("ini pilihan ke-3")
		case 4 : 
			fmt.Println("ini pilihan ke-4")
		default : 
			fmt.Printf("Salah")
	}

}
