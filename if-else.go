package main

import "fmt"

func main(){

	var a, b, c int
	fmt.Print("Введите число A: ")
	fmt.Scan(&a)
	fmt.Print("Введите число B: ")
	fmt.Scan(&b)
	fmt.Print("Введите число C: ")
	fmt.Scan(&c)


	if a > b && a != b{
		if a > c && a!= c {
			fmt.Println(b,"<",a,">",c)
		}else {
			fmt.Println(a,"<",c,">",b)
		}
	}


	if a > b && a != b {
		fmt.Println(a,">",b)
	} else if a == b {
		fmt.Println(a,"=",b)
	} else {
		fmt.Println(a,"<",b)
	}


	if b > c {
		fmt.Println(b,">",c)
	} else if b == c {
		fmt.Println(b,"=",c)
	} else {
		fmt.Println(b,"<",c)
	}


}
