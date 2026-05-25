package main
import (
	"fmt"
)
type City struct {
	Name       string  // Название города
	Population int     // Население (тыс. чел.)
	Country    string  // Страна
	Area       float64 // Площадь (кв. км)
}
func NewCity(name, country string, population int, area float64) *City {
	return &City{
		Name:       name,
		Country:    country,
		Population: population,
		Area:       area,
	}
}
func (c *City) PrintInfo() {
	fmt.Printf("Город: %s\nСтрана: %s\nНаселение: %d тыс. чел.\nПлощадь: %.2f кв. км\n",
		c.Name, c.Country, c.Population, c.Area)
	fmt.Println("-------------------------")
}
func (c *City) SetPopulation(newPopulation int) {
	if newPopulation >= 0 {
		oldPopulation := c.Population
		c.Population = newPopulation
		fmt.Printf("Население города %s обновлено: было %d тыс. чел., стало %d тыс. чел.\n",
			c.Name, oldPopulation, c.Population)
	} else {
		fmt.Printf("Ошибка: население не может быть отрицательным (%d)\n", newPopulation)
	}
	fmt.Println("-------------------------")
}
func (c *City) GetPopulationDensity() float64 {
	if c.Area <= 0 {
		fmt.Printf("Ошибка: площадь города %s не может быть равна нулю или отрицательной\n", c.Name)
		return 0
	}
	density := float64(c.Population*1000) / c.Area
	return density
}
func main() {
	city := NewCity("Москва", "Россия", 12500, 2511.0)
	fmt.Println("=== ИНФОРМАЦИЯ О ГОРОДЕ ===")
	city.PrintInfo()=
	density := city.GetPopulationDensity()
	fmt.Printf("Плотность населения: %.2f чел./кв. км\n", density)
	fmt.Println("-------------------------")
	fmt.Println("=== ОБНОВЛЕНИЕ ДАННЫХ ===")
	city.SetPopulation(13000)
	fmt.Println("=== ОБНОВЛЕННАЯ ИНФОРМАЦИЯ ===")
	city.PrintInfo()
	newDensity := city.GetPopulationDensity()
	fmt.Printf("Обновленная плотность населения: %.2f чел./кв. км\n", newDensity)
	fmt.Println("-------------------------")
}
