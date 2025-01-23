package main

import (
	"fmt"
)

// исходная конфигурация трёх роторов
var config0 = [3]rune{'c', 'f', 'w'}

// конфигурация, которая будет меняться при сдвигах
var config = [3]rune{'c', 'f', 'w'}

// проверка на то, является ли символ буквой латинского алфавита
func isLetter(a rune) bool {
	return (a >= 'a') && (a <= 'z')
}

// функция сдвига буквы по модулю 26
func add(a rune, b int) rune {
	if b < 0 {
		b = 26 + b
	}
	if b <= int('z'-a) {
		return a + rune(b)
	}
	return 'a' + rune(b) - ('z' - a + 1)
}

// функция преобразования для роторов
func rotor(n int, a rune, flag bool) rune {
	if flag {
		return add(a, int(config[n]-'a'))
	}
	return add(a, -int(config[n]-'a'))
}

// функция преобразования для рефлектора
// в данном случае - зеркалим относительно середины алфавита
func reflector(n rune) rune {
	return add('a', int('z'-n))
}

// функция сдвига для роторов
func shift() {
	if config[0] == 'z' {
		config[0] = 'a'
	} else {
		config[0]++
	}
	if config[0] == config0[0] {
		if config[1] == 'z' {
			config[1] = 'a'
		} else {
			config[1]++
		}
		if config[1] == config0[1] {
			if config[2] == 'z' {
				config[2] = 'a'
			} else {
				config[2]++
			}
		}
	}
}

func main() {
	var str string
	fmt.Print("Введите строку : ")
	fmt.Scan(&str)
	s := []rune(str)
	res := 'a'
	fmt.Print("Шифр для строки : ")
	for c := range s {
		res = rune(s[c])
		if !isLetter(res) {
			continue
		}
		for i := 0; i < 3; i++ {
			res = rotor(i, res, true)
		}
		res = reflector(res)
		for i := 2; i >= 0; i-- {
			res = rotor(i, res, false)
		}
		shift()
		fmt.Print(string(res))
	}
	fmt.Println()
}
