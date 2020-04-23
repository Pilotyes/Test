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
	for {
		x, err = readFloat("Введите координату X:")
		if err == nil {
			break
		} else {
			fmt.Println("НЕкорректное значение")
		}
	}
	for {
		y, err = readFloat("Введите координату Y:")
		if err == nil {
			break
		} else {
			fmt.Println("НЕкорректное значение")
		}
	}
	for {
		r, err = readFloat("Введитe radius R:")
		if err == nil {
			break
		} else {
			fmt.Println("НЕкорректное значение")
		}
	}
	if beBelong(x, y, r) == false {
		fmt.Printf("Tochka s coordinatami x = %f, y = %f - ne popadaet v krug s radiusom %f\n", x, y, r)
	} else {
		fmt.Printf("Tochka s coordinatami x = %f, y = %f - popadaet v krug s radiusom %f\n", x, y, r)
	}
}
