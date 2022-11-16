package main

import (
	"fmt"
	"io"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func insertionSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := i; j > 0 && arr[j-1] > arr[j]; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
	return arr
}

func Sort(a []int) []int {
	result := make([]int, len(a))
	copy(result, a)
	for i := 1; i < len(result); i++ {
		key := result[i]
		j := i - 1
		for j >= 0 && result[j] > key {
			result[j+1] = result[j]
			j--
		}
		result[j+1] = key
	}
	return result
}
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
	e := []int{100, 1000, 2000, 3000, 4000, 5000, 6000, 7000, 8000,
		9000, 10000, 20000, 30000, 40000, 50000}

	f, err := os.Create("tiemposejecucioninsert/tiemposgo.txt")
	if err != nil {
		fmt.Printf("error", err)
		return
	}

	f1, err1 := os.Create("desviacionestandarinsert/desviancionstandargolang.txt")
	if err1 != nil {
		fmt.Printf("error", err1)
		return
	}

	datos := "datos"

	defer f.Close()

	for i := 0; i < 15; i++ {
		var t []float64

		for j := 0; j < 5; j++ {
			rand.Seed(time.Now().Unix())

			lista := generarlistas(datos + strconv.Itoa(j+1) + ".txt")

			start := time.Now()
			Sort(lista[i])
			duration := time.Since(start).Seconds()
			t = append(t, duration)
		}

		var prom float64 = 0
		for k := 0; k < 5; k++ {
			prom += t[k]
		}

		var promf float64 = prom / 5

		var des float64 = 0
		for d := 0; d < 5; d++ {
			des += math.Pow(t[d]-promf, 2)
		}
		des = math.Sqrt(des / 4)

		fmt.Println("promedio array de tamaño ", e[i], ": ", promf, " Desviacion estandar: ", des)

		_, err = f.WriteString(fmt.Sprintf("%d\n", promf))
		if err != nil {
			fmt.Printf("error", err)
		}

		_, err1 = f1.WriteString(fmt.Sprintf("%d\n", des))
		if err1 != nil {
			fmt.Printf("error", err1)
		}
	}

}
