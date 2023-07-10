package main

var justString string

func someFunc() {
	/*
		утечка памяти
		panic если строка содержит менее 100 символов
	*/
	v := createHugeString(1 << 10)
	copyJustString(v)
}

func copyJustString(v string) {
	length := len(v)
	if length > 100 {
		length = 100
	}
	tmp := make([]byte, length)
	copy(tmp, v[:length])
	justString = string(tmp)
}

func createHugeString(a int) string {
	//...
	return ""
}

func main() {
	someFunc()
}
