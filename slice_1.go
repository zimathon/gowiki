package main

import {
	"fmt"
}

function main(){
	var a [4]int

	a[0]=1

	i := a[0]

	fmt.Printf("a[0]=%d\n", a[0])
	fmt.Printf("i=%d\n", i)

	if i == a[0]{
		fmt.Println("i == a[0]")
	}else{
		fmt.Println("i != a[0])
	}

	for j := 0; j < len(a); j++{
		if a[j] == 0{
			fmt.Printf("a[%d] == %d\n", j, a[j])
		}else{
			fmt.Printf("a[%d] != %d\n", j, a[j])
		}
	}
}

