package main

import (
	"fmt"
	"io"
	"os"
)

func generarlistas(x string) [][]int {

	e := []int{100, 1000, 2000, 3000, 4000, 5000, 6000, 7000, 8000,
		9000, 10000, 20000, 30000, 40000, 50000}

	file, err := os.Open(x)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var perline int
	var nums []int
	var listadatos [][]int
	for {

		_, err := fmt.Fscanf(file, "%d\n", &perline) // give a patter to scan

		if err != nil {

			if err == io.EOF {
				break // stop reading the file
			}
			fmt.Println(err)
			os.Exit(1)
		}

		nums = append(nums, perline)
	}

	for i := 0; i < len(e); i++ {
		temp := nums[0:e[i]]
		listadatos = append(listadatos, temp)
	}

	return listadatos
}

func main() {
	lista := generarlistas("datos1.txt")
	fmt.Println(lista[0])

	lista1 := generarlistas("datos2.txt")
	fmt.Println(lista1[0])
}
