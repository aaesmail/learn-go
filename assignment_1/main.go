package main

import "fmt"

func main() {
	list := []int{}

	for i := 0; i <= 10; i++ {
		list = append(list, i)
	}

	for _, item := range list {
		if item % 2 == 0 {
			fmt.Println(fmt.Sprint(item), "is even")
		} else {
			fmt.Println(fmt.Sprint(item), "is odd")
		}
	}
}
