package main

import (
	"fmt"
	"sync"
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

	y := make([][]float64, n)
	for i := range a {
		y[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			y[i][j] = 0.0
		}
	}

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

	var wg sync.WaitGroup
	wg.Add(n)

	work_forward := func(i int, j int) {
		defer wg.Done()
		if j <= i {
			return
		}
		y[i][j] = a[j][i] / a[i][i]
		for k := i; k < n; k++ {
			a[j][k] -= (a[i][k] * y[i][j])
		}
		b[j] -= (b[i] * y[i][j])
	}

	work_back := func(i int, j int) {
		defer wg.Done()
		if j >= i {
			return
		}
		y[i][j] = a[j][i]
		for k := 0; k < n; k++ {
			a[j][k] -= (a[i][k] * y[i][j])
		}
		b[j] -= (b[i] * y[i][j])
	}

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

		for j := 0; j < n; j++ {
			go work_forward(i, j)
		}
		wg.Wait()
		wg.Add(n)

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
		for j := 0; j < n; j++ {
			go work_back(i, j)
		}
		wg.Wait()
		wg.Add(n)

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
