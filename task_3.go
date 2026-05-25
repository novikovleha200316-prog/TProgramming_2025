package main

import (
	"fmt"
	"math"
)
func calculateY(a, x float64) (float64, error) {
	xSquared := x * x
	lgX2 := math.Log10(xSquared)
	if lgX2-a*a <= 0 {
		return 0, fmt.Errorf("подкоренное выражение отрицательное или ноль (lg(x^2)=%.4f, a^2=%.4f)", lgX2, a*a)
	}
	numerator := a - lgX2
	denominator := math.Sqrt(lgX2 - a*a)
	y := numerator / denominator
	return y, nil
}
func main() {
	a := 0.1
	fmt.Println("=== Задача А (вычисление для набора значений) ===")
	xValuesA := []float64{0.15, 1.37, 0.25, 0.2, 0.3, 0.44, 0.6, 0.56}
	for _, x := range xValuesA {
		y, err := calculateY(a, x)
		if err != nil {
			fmt.Printf("При x = %.2f \t ОШИБКА: %s\n", x, err)
		} else {
			fmt.Printf("При x = %.2f \t y = %.4f\n", x, y)
		}
	}
	fmt.Println("\n=========================================\n")
	fmt.Println("=== Задача В (массив значений) ===")
	xValuesB := []float64{0.15, 0.2, 0.3, 0.44, 0.56, 0.6} 
	for _, x := range xValuesB {
		y, err := calculateY(a, x)
		if err != nil {
			fmt.Printf("При x = %.2f \t ОШИБКА: %s\n", x, err)
		} else {
			fmt.Printf("При x = %.2f \t y = %.4f\n", x, y)
		}
	}
}
