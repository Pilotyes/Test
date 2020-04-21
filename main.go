/*Version 2:
Расширить функционал программы так, чтобы координаты точки и радиус круга вводились пользователем с клавиатуры. Предусмотреть возможность ввода некорректных данных. Диалог с пользователем:
"Введите координату X: "
"Введите координату Y: "
"Введите радиус R: "
Закоммитить вторую версию кода
*/

package main

import "fmt"

func main() {
	var x, y, r float64
	fmt.Println("Введите координату X:")
	fmt.Scanln(&x)
	fmt.Println("Введите координату Y:")
	fmt.Scanln(&y)
	fmt.Println("Введите радиус R:")
	fmt.Scanln(&r)
	if (x*x + y*y) > r*r {
		fmt.Printf("Tochka s coordinatami x = %f, y = %f - ne popadaet v krug s radiusom %f", x, y, r)
	} else {
		fmt.Printf("Tochka s coordinatami x = %f, y = %f - popadaet v krug s radiusom %f", x, y, r)
	}
}
