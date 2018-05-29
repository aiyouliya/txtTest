package main

import "fmt"

func main(){
	sum :=1
	i :=0
	for sum <1000{
		sum+= sum
		i+=1;
		fmt.Println("i = ",i,"sum = ",sum)
		
	}
	fmt.Println(sum)
	fmt.Println(i)
}