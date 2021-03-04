package main 

import "fmt"

func main(){

	// uint32 : 0 sampai 4294967295
	// uint64 : 0 sampai 18446744073709551615

	decimalNumber := 2.62
	decimalNumber2 := 2.622

	fmt.Printf("Angka : %f\n", decimalNumber)
	fmt.Printf("Angka : %.2f\n", decimalNumber)
	fmt.Printf("Angka : %f\n", decimalNumber2)
	fmt.Printf("Angka : %.3f", decimalNumber2)


}