package main

import (
	"fmt"
	"math"
)

// Функция для вычисления значения y
func calculateY(a, b, x float64) (float64, error) {
	if x <= 0 {
		return 0, fmt.Errorf("недопустимое значение x")
	}
	if x == 1 {
		return 0, fmt.Errorf("логарифм нуля не определен")
	}
	if x == 2 {
		return 0, fmt.Errorf("деление на ноль (lg|2-1| = 0)")
	}

	log5x := math.Log(x) / math.Log(5)
	numerator := a*math.Sqrt(x) - b*log5x
	denominator := math.Log10(math.Abs(x - 1))

	return numerator / denominator, nil
}

func main() {
	a := 4.1
	b := 2.7

	// ==========================================
	// ЗАДАЧА А
	// ==========================================
	fmt.Println("=== Задача А (вычисление через массив) ===")
	xStart := 1.2
	xEnd := 5.2
	step := 0.8

	// 1. Создаем пустой массив (срез) для хранения значений x
	var arrayA []float64

	// Заполняем массив значениями от xStart до xEnd с шагом step
	for x := xStart; x <= xEnd+0.0001; x += step {
		arrayA = append(arrayA, x)
	}

	// 2. Используем цикл for для прохождения по сформированному массиву
	for i := 0; i < len(arrayA); i++ {
		x := arrayA[i] // Берем элемент массива по индексу i
		y, err := calculateY(a, b, x)
		if err != nil {
			fmt.Printf("При x = %.2f \t ОШИБКА: %s\n", x, err)
		} else {
			fmt.Printf("При x = %.2f \t y = %.4f\n", x, y)
		}
	}

	fmt.Println("\n==================================================\n")

	// ==========================================
	// ЗАДАЧА В
	// ==========================================
	fmt.Println("=== Задача В (вычисление через массив) ===")

	// 1. Сразу инициализируем массив заданными значениями
	arrayB := []float64{1.9, 2.15, 2.34, 2.73, 3.16}

	// 2. Используем классический цикл for для прохождения по массиву
	for i := 0; i < len(arrayB); i++ {
		x := arrayB[i] // Берем элемент массива по индексу i
		y, err := calculateY(a, b, x)
		if err != nil {
			fmt.Printf("При x = %.2f \t ОШИБКА: %s\n", x, err)
		} else {
			fmt.Printf("При x = %.2f \t y = %.4f\n", x, y)
		}
	}
}
