/*Version 3:
Вынести вычисление принадлежности точки с координатами в отдельную функцию с семантикой:
Входящие аргументы:
	x, y, r: float
Возвращяемые значения:
	ret: bool
Закоммитить третью версию кода
*/

package main

import "fmt"

func beBelong(x, y, r float64) bool {
	return (x*x + y*y) <= r*r
}

func main() {
	var x, y, r float64
	for {
		fmt.Println("Введите координату X:")
		_, err := fmt.Scanf("%f", &x)
		if err == nil {
			break
		} else {
			fmt.Println("НЕкорректное значение")
		}
	}
	fmt.Println("Введите координату Y:")
	fmt.Scanln(&y)
	fmt.Println("Введите радиус R:")
	fmt.Scanln(&r)
	if beBelong(x, y, r) == false {
		fmt.Printf("Tochka s coordinatami x = %f, y = %f - ne popadaet v krug s radiusom %f\n", x, y, r)
	} else {
		fmt.Printf("Tochka s coordinatami x = %f, y = %f - popadaet v krug s radiusom %f\n", x, y, r)
	}
}
