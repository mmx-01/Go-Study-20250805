package common

// 加法
func Sum(a, b int) int {
	return a + b
}

// 减法
func Subtraction(a, b int) int {
	return a - b
}

// 乘法
func Dultiplication(a, b int) int {
	return a * b
}

// 除法
func Division(a, b int) int {
	return a / b
}

// 选择计算类型 + - * /
func Search(a string) string {
	switch a {
	case "1":
		return "加法"
	case "2":
		return "减法"
	case "3":
		return "乘法"
	case "4":
		return "除法"
	}
	return "0"
}
