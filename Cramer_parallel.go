package main

import (
	"fmt"
	"log"
	"math"
	"sync"
	"time"
)

// создаём тип данных для матриц
type Matrix [][]float64

func (m Matrix) Rows() int {
	return len(m)
}

func (m Matrix) Columns() int {
	return len(m[0])
}

func (m Matrix) IsSquare() bool {
	return m.Columns() == m.Rows()
}

func (m Matrix) IsMatrix() bool {
	if m.Rows() == 0 {
		return false
	}

	for _, row := range m {
		if len(m[0]) != len(row) {
			return false
		}
	}
	return true
}

func InBetween(i, min, max int) bool {
	if (i >= min) && (i <= max) {
		return true
	} else {
		return false
	}
}

func (m Matrix) ExcludeColumn(col_index int) (Matrix, error) {

	if !InBetween(col_index, 1, m.Columns()) {
		return Matrix{}, fmt.Errorf("input not in range")
	}

	result := make(Matrix, m.Rows())
	for i, row := range m {
		for j, el := range row {
			if j == col_index-1 {
				continue
			}
			result[i] = append(result[i], el)
		}
	}
	return result, nil
}

func (m Matrix) ExcludeRow(row_index int) (Matrix, error) {
	if !InBetween(row_index, 1, m.Rows()) {
		return Matrix{}, fmt.Errorf("input not in range")
	}

	var result Matrix
	for i, r := range m {
		if i == row_index-1 {
			continue
		}
		result = append(result, r)
	}
	return result, nil
}

func (m Matrix) Det() (float64, error) {

	if !m.IsMatrix() || !m.IsSquare() {
		return -1, fmt.Errorf("determinant is not defined for the input [Matrix: %t][Square: %t]",
			m.IsMatrix(), m.IsSquare())
	}

	if m.Rows() == 1 {
		return m[0][0], nil
	}

	if m.Rows() == 2 {
		return m[0][0]*m[1][1] - m[0][1]*m[1][0], nil
	}

	// исключаем первый столбец
	partial_matrix, err := m.ExcludeRow(1)
	if err != nil {
		return -1, err
	}

	var temp float64 = 0

	// раскладываем по элементам первого столбца
	for i, el := range m[0] {

		reduced_matrix, err := partial_matrix.ExcludeColumn(i + 1)
		if err != nil {
			return -1, err
		}

		partial_det, err := reduced_matrix.Det()
		if err != nil {
			return -1, err
		}

		temp = temp + partial_det*el*math.Pow(-1, float64(i))
	}

	return temp, nil
}

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

	x := Matrix(a)
	det, err := Matrix(a).Det()
	if err != nil {
		log.Fatalf("Error in calculating the determinant: %v", err)
	}
	fmt.Println(det)
	for i := 0; i < x.Rows(); i++ {
		for j := 0; j < x.Columns(); j++ {
			fmt.Printf("%f ", x[i][j])
		}
		fmt.Println()
	}
	mainDet := det

	fmt.Print("Введите столбец : ")

	b := make([]float64, n)
	for i := range b {
		fmt.Scan(&b[i])
	}

	start := time.Now()

	res := make([]float64, n)

	dets := make([][][]float64, n)
	for i := range dets {
		dets[i] = make([][]float64, n)
		for j := 0; j < n; j++ {
			dets[i][j] = make([]float64, n)
			for k := 0; k < n; k++ {
				dets[i][j][k] = a[j][k]
			}
		}
	}

	for i := 0; i < n; i++ {
		for k := 0; k < n; k++ {
			dets[i][k][i] = b[k]
		}
	}

	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				fmt.Print(dets[k][i][j])
				fmt.Print(" ")
			}
			fmt.Println()
		}
	}

	if mainDet == 0 {
		fmt.Printf("Матрица вырожденна")
		return
	}

	var wg sync.WaitGroup
	wg.Add(n)
	work := func(i int) {
		defer wg.Done()
		det, err = Matrix(dets[i]).Det()
		if err != nil {
			log.Fatalf("Error in calculating the determinant: %v", err)
		}
		fmt.Println(det)
		res[i] = det / mainDet
	}

	for i := 0; i < n; i++ {
		go work(i)
	}
	wg.Wait()

	for i := 0; i < n; i++ {
		fmt.Printf("x_%d = %g \n", i+1, res[i])
	}
	fmt.Printf("Время работы программы : %v мкс", time.Now().Sub(start).Microseconds())
	fmt.Println()
}
