package main

import (
	"fmt"
	"math"
)

// Функция для вычисления значения y по заданной формуле
func calculateY(a, b, x float64) (float64, error) {
	// Проверка области определения
	if x <= 0 {
		return 0, fmt.Errorf("недопустимое значение x (корень и логарифм отрицательного числа)")
	}
	if x == 1 {
		return 0, fmt.Errorf("логарифм нуля не определен (x = 1)")
	}
	if x == 2 {
		return 0, fmt.Errorf("деление на ноль (lg|2-1| = 0)")
	}

	// Вычисляем числитель: a * sqrt(x) - b * log5(x)
	// log5(x) вычисляется как ln(x) / ln(5)
	log5x := math.Log(x) / math.Log(5)
	numerator := a*math.Sqrt(x) - b*log5x

	// Вычисляем знаменатель: lg(|x - 1|)
	denominator := math.Log10(math.Abs(x - 1))

	// Итоговое значение
	y := numerator / denominator
	return y, nil
}

func main() {
	// Исходные константы
	a := 4.1
	b := 2.7

	// ==========================================
	// ЗАДАЧА А: Табулирование функции
	// ==========================================
	fmt.Println("=== Задача А (вычисление с шагом) ===")
	xStart := 1.2
	xEnd := 5.2
	step := 0.8

	// Используем небольшую прибавку (0.0001) к xEnd для компенсации погрешностей float64
	for x := xStart; x <= xEnd+0.0001; x += step {
		y, err := calculateY(a, b, x)
		if err != nil {
			fmt.Printf("При x = %.2f \t ОШИБКА: %s\n", x, err)
		} else {
			fmt.Printf("При x = %.2f \t y = %.4f\n", x, y)
		}
	}

	fmt.Println("\n=========================================\n")

	// ==========================================
	// ЗАДАЧА В: Вычисление для конкретных значений
	// ==========================================
	fmt.Println("=== Задача В (массив значений) ===")
	xValues := []float64{1.9, 2.15, 2.34, 2.73, 3.16}

	for _, x := range xValues {
		y, err := calculateY(a, b, x)
		if err != nil {
			fmt.Printf("При x = %.2f \t ОШИБКА: %s\n", x, err)
		} else {
			fmt.Printf("При x = %.2f \t y = %.4f\n", x, y)
		}
	}
}
