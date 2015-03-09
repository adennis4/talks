package main

import (
	"fmt"
	"runtime"
)

type UselessStruct struct {
	MyInt int
}

func useless(args ...string) {
	runtime.Breakpoint()
	fmt.Println(args)
}

func main() {
	us := UselessStruct{MyInt: 0}
	greeting := "Hello ChicaGoLang"
	for {
		us.MyInt++
		if us.MyInt == 3 {
			break
		}
		useless(greeting, "is the", "best meetup")
	}
}
