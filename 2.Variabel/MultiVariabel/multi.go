package main 

import "fmt"

func main(){

	// multi variabel dengan manifest typing
	var first, second, third string 
	first, second, third = "satu","dua","tiga"

	var fourth, fifth, sixth string = "empat","lima","enam"

	// multi variabel dengan inference type

	var satu, dua, tiga = 1, 2, 3

	empat, lima, enam := 4,5,6

	fmt.Println(satu,dua,tiga,empat,lima,enam)
	fmt.Println(first,second,third,fourth,fifth,sixth)

}