package unzipString

import "strings"

/*
Задача на распаковку
Создать Go-функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы/руны, например:
"a4bc2d5e" => "aaaabccddddde"
"abcd" => "abcd"
"45" => "" (некорректная строка)
"" => ""

Дополнительно
Реализовать поддержку escape-последовательностей.
Например:
qwe\4\5 => qwe45 (*)
qwe\45 => qwe44444 (*)
qwe\\5 => qwe\\\\\ (*)


В случае если была передана некорректная строка, функция должна возвращать ошибку. Написать unit-тесты.
*/

//const (
//	numbers = "123456789"
//	letters = "abcdefghijklmnopqrstuvwxyz"
//)

// 48 - 57
//var chToNum map[byte]uint16 = map[byte]uint16{
//	48: 0,
//	49: 1,
//	50: 2,
//	51: 3,
//	52: 4,
//	53: 5,
//	54: 6,
//	55: 7,
//	56: 8,
//	57: 9,
//}

func multiplyChar(ch byte, n int) string {
	var builder strings.Builder
	for i := 0; i < n; i++ {
		builder.WriteByte(ch)
	}
	return builder.String()
}

func isLetter(ch byte) bool {
	if (64 < ch && ch < 91) || (96 < ch && ch < 123) {
		return true
	}
	return false
}

func isDigit(ch byte) bool {
	if 47 < ch && ch < 58 {
		return true
	}
	return false
}

func chToDigit(ch byte) int {
	return int(ch) - 48
}

func UnzipString(str string) string {
	if str == "" {
		return ""
	}

	var builder strings.Builder
	builder.Grow(len(str))

	for i := 0; i < len(str)-1; i++ {
		if isLetter(str[i]) {
			if isDigit(str[i+1]) {
				//for
				builder.WriteString(multiplyChar(str[i], chToDigit(str[i+1])))
				//i++
			} else {
				builder.WriteByte(str[i])
			}
		}
	}

	if !isDigit(str[len(str)-1]) {
		builder.WriteByte(str[len(str)-1])
	}
	return builder.String()
}
