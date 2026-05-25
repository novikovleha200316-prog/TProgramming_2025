package main
import (
	"fmt"
	"math"
)
func calculateY(a, b, x float64) (float64, error) {
	bx := b * x
	ax := a * x
	if math.Abs(ax) > 1 {
		return 0, fmt.Errorf("arccos(ax) не определен: |ax| = %.2f > 1", math.Abs(ax))
	}
	arccosAx := math.Acos(ax)
	denominator := b + arccosAx*arccosAx
	if math.Abs(denominator) < 1e-10 {
		return 0, fmt.Errorf("деление на ноль: b + arccos²(ax) = 0")
	}
	if bx <= 0 {
		return 0, fmt.Errorf("lg(bx) не определен: bx = %.2f <= 0", bx)
	}
	lgBx := math.Log10(bx)
	squaredLg := lgBx * lgBx
	numerator := a - squaredLg
	y := numerator / denominator
	return y, nil
}
func main() {
	a := 0.1
	b := 0.5
	fmt.Println("=== Задача А (вычисление через массив) ===")
	xValuesA := []float64{0.15, 1.37, 0.25, 0.2, 0.3, 0.44, 0.6, 0.56}
	for i := 0; i < len(xValuesA); i++ {
		x := xValuesA[i]
		y, err := calculateY(a, b, x)
		if err != nil {
			fmt.Printf("При x = %.2f \t ОШИБКА: %s\n", x, err)
		} else {
			fmt.Printf("При x = %.2f \t y = %.4f\n", x, y)
		}
	}
	fmt.Println("\n==================================================\n")
	fmt.Println("=== Задача В (вычисление через массив) ===")
	xValuesB := []float64{0.15, 0.2, 0.3, 0.44, 0.56, 0.6}
	for i := 0; i < len(xValuesB); i++ {
		x := xValuesB[i]
		y, err := calculateY(a, b, x)
		if err != nil {
			fmt.Printf("При x = %.2f \t ОШИБКА: %s\n", x, err)
		} else {
			fmt.Printf("При x = %.2f \t y = %.4f\n", x, y)
		}
	}
}
