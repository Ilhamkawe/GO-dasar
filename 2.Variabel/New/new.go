package main 

import "fmt"

func main(){

	Name := new(string) //pointer alokasi memori
	
	fmt.Println(Name)
	fmt.Println(*Name)


}