package main

import (
	"fmt"
	"time"
)

func main() {
	var n int
	fmt.Print("Введите размер матрицы n : ")
	fmt.Scan(&n)

	fmt.Print("Введите элементы матрицы : ")

	a := make([][]float64, n)
	for i := range a {
		a[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&a[i][j])
		}
	}

	fmt.Print("Введите столбец : ")

	b := make([]float64, n)
	for i := range b {
		fmt.Scan(&b[i])
	}

	start := time.Now()

	index := 0
	t := 0.0

	/*
		for k := 0; k < n; k++ {
			for j := 0; j < n; j++ {
				fmt.Print(a[k][j])
				fmt.Print(" ")
			}
			fmt.Print(" | ")
			fmt.Print(b[k])
			fmt.Println()
		}*/

	fmt.Println()

	//fmt.Println("Прямой ход")

	fmt.Println()

	// Прямой ход
	for i := 0; i < n; i++ {
		index = i
		for j := i; j < n; j++ {
			if a[j][i] != 0 {
				break
			}
			index += 1
		}

		if index == n {
			fmt.Println("Матрица вырожденна")
			return
		}

		for j := 0; j < n; j++ {
			t = a[index][j]
			a[index][j] = a[i][j]
			a[i][j] = t
		}
		t = b[index]
		b[index] = b[i]
		b[i] = t

		/*
			for k := 0; k < n; k++ {
				for j := 0; j < n; j++ {
					fmt.Print(a[k][j])
					fmt.Print(" ")
				}
				fmt.Print(" | ")
				fmt.Print(b[k])
				fmt.Println()
			}*/

		fmt.Println()

		for j := i + 1; j < n; j++ {
			t = a[j][i] / a[i][i]
			for k := i; k < n; k++ {
				a[j][k] -= (a[i][k] * t)
			}
			b[j] -= (b[i] * t)
		}

		/*
			for k := 0; k < n; k++ {
				for j := 0; j < n; j++ {
					fmt.Print(a[k][j])
					fmt.Print(" ")
				}
				fmt.Print(" | ")
				fmt.Print(b[k])
				fmt.Println()
			}*/
	}

	fmt.Println()

	//fmt.Println("Обратный ход")

	fmt.Println()

	// Обратный ход

	for i := n - 1; i >= 0; i-- {
		b[i] /= a[i][i]
		a[i][i] = 1
		for j := 0; j < i; j++ {
			t = a[j][i]
			for k := 0; k < n; k++ {
				a[j][k] -= (a[i][k] * t)
			}
			b[j] -= (b[i] * t)
		}

		/*
			for k := 0; k < n; k++ {
				for j := 0; j < n; j++ {
					fmt.Print(a[k][j])
					fmt.Print(" ")
				}
				fmt.Print(" | ")
				fmt.Print(b[k])
				fmt.Println()
			}*/

		fmt.Println()
	}
	for i := 0; i < n; i++ {
		fmt.Printf("x_%d = %g \n", i+1, b[i])
	}
	fmt.Printf("Время работы программы : %v мкс", time.Now().Sub(start).Microseconds())
	fmt.Println()
}
