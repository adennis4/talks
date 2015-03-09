package main

import "fmt"

type UselessStruct struct {
	MyInt int
}

func main() {
	us := UselessStruct{MyInt: 0}
	greeting := "Hello ChicaGoLang"
	for {
		us.MyInt++
		if us.MyInt == 3 {
			break
		}
		fmt.Println(greeting)
	}
}
