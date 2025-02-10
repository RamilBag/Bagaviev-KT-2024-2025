package main

import (
	"fmt"
	"io"
	"os"
	"unicode"
)

// исходная конфигурация трёх роторов
var config0 = [3]rune{'c', 'f', 'w'}

// конфигурация, которая будет меняться при сдвигах
var config = [3]rune{'c', 'f', 'w'}

// проверка на то, является ли символ (строчной) буквой латинского алфавита
func isLowerLetter(a rune) bool {
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

	res := 'a'
	ans := ""
	file, err := os.Open("/Users/ramilbagaviev/Downloads/Golang/Enigma/readFile/input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	data := make([]byte, 64)

	for {
		n, err := file.Read(data)
		if err == io.EOF { // если конец файла
			break // выходим из цикла
		}
		str += string(data[:n])
	}

	s := []rune(str)
	isLower := false

	ffile, errr := os.OpenFile("/Users/ramilbagaviev/Downloads/Golang/Enigma/readFile/output.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if errr != nil {
		fmt.Println("Unable to open file:", errr)
		os.Exit(1)
	}
	defer ffile.Close()

	for c := range s {
		res = rune(s[c])
		isLower = true
		if !isLowerLetter(res) {
			if !isLowerLetter(unicode.ToLower(res)) {
				continue
			}
			isLower = false
			res = unicode.ToLower(res)
		}
		for i := 0; i < 3; i++ {
			res = rotor(i, res, true)
		}
		res = reflector(res)
		for i := 2; i >= 0; i-- {
			res = rotor(i, res, false)
		}
		shift()
		if !isLower {
			res = unicode.ToUpper(res)
		}
		fmt.Print(string(res))
		ans += string(res)
	}
	fmt.Println()
	ffile.WriteString(ans)
}

