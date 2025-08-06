package pkg1

func Demo1() string {
	var _a = "外部包，打印外部变量"
	return "内部包，只写本项目可访问的文件\n" + _a
	//return 1 + 1
}
