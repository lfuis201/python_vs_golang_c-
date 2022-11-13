package main
import (
    "fmt"
    "math/rand"
    "time"
	"os"
	"math"
)

func bubblesort(array[] int)[]int {
   for i:=0; i< len(array)-1; i++ {
      for j:=0; j < len(array)-i-1; j++ {
         if (array[j] > array[j+1]) {
            array[j], array[j+1] = array[j+1], array[j]
         }
      }
   }
   return array
}

func main() {
	e := []int {100,1000, 2000, 3000, 4000, 5000, 6000, 7000, 8000,
	9000, 10000, 20000, 30000, 40000, 50000}

	f, err :=os.Create("tiemposgo.txt")
	if err != nil {
		fmt.Printf("error", err)
		return
	}

	f1, err1 :=os.Create("desviancionstandargolang.txt")
	if err1 != nil {
		fmt.Printf("error", err1)
		return
	}

	defer f.Close()

	for i:=0; i< 15; i++ {
		var t []float64

		for j:=0; j< 5; j++{
			rand.Seed(time.Now().Unix())

    		//Generate a random array of length n
    		var n int = e[i]
    		var a []int
    		a = rand.Perm(n)
    		start := time.Now()
    		bubblesort(a)
    		duration := time.Since(start).Seconds()
			t = append(t, duration)
		}

		var prom float64 = 0
		for k:=0; k< 5; k++{
			prom += t[k]
		}

		var promf float64 = prom/5

		var des float64 = 0
		for d:=0; d< 5; d++{
			des += math.Pow(t[d]-promf, 2)
		}
		des = math.Sqrt(des/4)
			
		fmt.Println("promedio array de tamaño ",e[i], ": ",promf," Desviacion estandar: ", des)

		_, err = f.WriteString(fmt.Sprintf("%d\n", promf))
		if err != nil{
			fmt.Printf("error", err)
		}

		_, err1 = f1.WriteString(fmt.Sprintf("%d\n", des))
		if err1 != nil{
			fmt.Printf("error", err1)
		}
	}

}