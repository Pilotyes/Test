/*Создать программу, которая вычисляет, попадает ли точка с захардкоженными координатами x, y в площадь круга с захардкоженным радиусом r.
Результат вывести в формате "Точка с координатами X=x, Y=y попадает|не попадает в круг с радиусом r"
Закоммитить первую версию кода
*/

package main

import "fmt"

func main() {
	var x, y, r float64
	x = 0.5
	y = 1
	r = 0.2
	if (x*x + y*y) > r {
		fmt.Printf("Tochka s coordinatami x = %f, y = %f - ne popadaet v krug s radiusom %f", x, y, r)
	} else {
		fmt.Printf("Tochka s coordinatami x = %f, y = %f - popadaet v krug s radiusom %f", x, y, r)
	}
}
