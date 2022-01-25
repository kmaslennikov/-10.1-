package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Введите значение X:")
	var x float64
	fmt.Scan(&x)

	fmt.Println("Введите точность:")
	var accuracy float64
	fmt.Scan(&accuracy)

	epsilon := math.Pow(10, (-1 * accuracy))
	var ex, n float64

	factorial := 1.0
	preview := 1.0
	for n = 0; ; n++ {
		factorial = factorial * (n + 1)		
		ex = preview + math.Pow(x, float64(n))/factorial

		if math.Abs(ex-preview) < epsilon {
			break
		}
		preview = ex	
	}
	fmt.Printf("Для x = %v и n = %v ex = %v\n", x, accuracy, ex)
}
