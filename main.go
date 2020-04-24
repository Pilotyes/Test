/*Version 3:
Вынести вычисление принадлежности точки с координатами в отдельную функцию с семантикой:
Входящие аргументы:
	x, y, r: float
Возвращяемые значения:
	ret: bool
Закоммитить третью версию кода
*/

package main

import (
	"fmt"
	"math"
	"strconv"
)

func beBelong(x, y, r float64) bool {
	return (math.Pow(x, 2.0) + math.Pow(y, float64(2))) <= math.Pow(r, 2.0)
}
func readFloat(str string) (float64, error) {
	fmt.Println(str)
	var tmpStr string
	_, err := fmt.Scanln(&tmpStr)
	if err != nil {
		return float64(0), err
	}
	ret, err := strconv.ParseFloat(tmpStr, 64)
	if err != nil {
		return float64(0), err
	}
	return ret, nil
}
func main() {
	var x, y, r float64
	var err error

	questions := []string{
		"Введите координату X:",
		"Введите координату Y:",
		"Введитe радиус R:",
	}

	for index, r := range []*float64{&x, &y, &r} {
		for {
			*r, err = readFloat(questions[index])
			if err == nil {
				break
			} else {
				fmt.Println("Некорректное значение")
			}
		}
	}
	if beBelong(x, y, r) == false {
		fmt.Printf("Точка с координатами X = %f, Y = %f - не попадает в круг с радиусом %f\n", x, y, r)
	} else {
		fmt.Printf("Точка с координатами X = %f, Y = %f - попадает в круг с радиусом %f\n", x, y, r)
	}
}
