package main

import "fmt"

func main() {
	desafio1()
	desafio2()
}

func desafio1() {
	for i := 0; i < 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}
}

func desafio2() {
	for i := 0; i < 100; i++ {
		if i%3 != 0 && i%5 != 0 {
			fmt.Println(i)
		} else {
			if i%3 == 0 {
				fmt.Println("Pin ")
			}

			if i%5 == 0 {
				fmt.Println("Pan ")
			}
		}
	}
}
