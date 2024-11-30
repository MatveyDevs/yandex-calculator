package calculation

import (
	"strconv"
	"strings"
)

func stringToFloat64(str string) float64 {
	degree := float64(1)
	var res float64 = 0
	var inverse = false
	for i := len(str); i > 0; i-- {
		if str[i-1] == '-' {
			inverse = true
		} else {
			res += float64(9-int('9'-str[i-1])) * degree
			degree *= 10
		}
	}
	if inverse {
		res = 0 - res
	}
	return res
}

func isSign(value rune) bool {
	return value == '+' || value == '-' || value == '*' || value == '/'
}

func Calc(expression string) (float64, error) {
	if len(expression) < 3 {
		return 0, ErrInvalidExpression
	}
	var res float64
	var b string
	var c rune = 0
	var resflag = false
	var isc int
	var countc = 0
	for _, value := range expression {
		if isSign(value) {
			countc++
		}
	}
	if isSign(rune(expression[0])) || isSign(rune(expression[len(expression)-1])) {
		return 0, ErrInvalidCharacter
	}
	for i, value := range expression {
		if value == '(' {
			isc = i
		}
		if value == ')' {
			calc, err := Calc(expression[isc+1 : i])
			if err != nil {
				return 0, ErrInvalidExpression
			}
			calcStr := strconv.FormatFloat(calc, 'f', 0, 64)
			i2 := i
			i -= len(expression[isc:i+1]) - len(calcStr)
			expression = strings.Replace(expression, expression[isc:i2+1], calcStr, 1)
		}
	}
	if countc > 1 {
		for i := 1; i < len(expression); i++ {
			value := rune(expression[i])
			if value == '*' || value == '/' {
				var imin = i - 1
				if imin != 0 {
					for !isSign(rune(expression[imin])) && imin > 0 {
						imin--
					}
					imin++
				}
				var imax = i + 1
				if imax == len(expression) {
					imax--
				} else {
					for !isSign(rune(expression[imax])) && imax < len(expression)-1 {
						imax++
					}
				}
				if imax == len(expression)-1 {
					imax++
				}
				calc, err := Calc(expression[imin:imax])
				if err != nil {
					return 0, ErrInvalidExpression
				}
				calcstr := strconv.FormatFloat(calc, 'f', 0, 64)
				i -= len(expression[isc:i+1]) - len(calcstr) - 1
				expression = strings.Replace(expression, expression[imin:imax], calcstr, 1)
			}
			if value == '+' || value == '-' || value == '*' || value == '/' {
				c = value
			}
		}
	}
	for _, value := range expression + "s" {
		switch {
		case value == ' ':
			continue
		case value > 47 && value < 58:
			b += string(value)
		case isSign(value) || value == 's':
			if resflag {
				switch c {
				case '+':
					res += stringToFloat64(b)
				case '-':
					res -= stringToFloat64(b)
				case '*':
					res *= stringToFloat64(b)
				case '/':
					res /= stringToFloat64(b)
				}
			} else {
				resflag = true
				res = stringToFloat64(b)
			}
			b = strings.ReplaceAll(b, b, "")
			c = value
		default:
			return 0, ErrDivisionByZero
		}
	}
	return res, nil
}
