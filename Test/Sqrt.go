package main

import (
	"fmt"
)

// Вычисление Квадратного корня
func Sqrt(x float64) float64 {
	if x <= 0{
		return 0
	}
	z := float64(1)
	tmp := z
	tmptmp := tmp
	for {
		z -= (z*z - x) / (2 * z) // точность
		if tmp == z {
			return z
		} else {
			if tmptmp == z {
				return z
			}
		}
		tmptmp = tmp
		tmp = z
	}
}

func main() {
	fmt.Println(Sqrt(2))
}
