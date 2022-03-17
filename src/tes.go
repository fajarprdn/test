package main

import "fmt"

func main() {
	var name string = "Fajar"
	fmt.Println(name)

	var num int = 432
	for i := 0; i <= num; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		} else {
			fmt.Println("a")
		}

	}
	var city = [4]string{
		"Jakarta",
		"Hanoi",
		"Kuala Lumpur",
	}
}
