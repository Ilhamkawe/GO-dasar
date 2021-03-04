package main 

import "fmt"

func main(){

	// FOR
	// for argumen { statement }

	for i := 0; i < 10; i++{

		fmt.Println("angka",i)

	}

	//FOR dengan argumen hanya kondisi
	//for kondisi { statement }
	i := 0
	for i < 10 {
		fmt.Println("angka",i)
		i++
	}

	// For ranpa argumen (infinity loop)
	// for{stateement}

	for {
		fmt.Println("halo")
	}



}