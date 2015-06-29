package main

import "fmt"

func convertToTheNumber10(i int) (j int) {
	defer func() {
		j /= 2
	}()
	j = 10
	return
}

func main() {
	i := convertToTheNumber10(5)
	if i == 10 {
		fmt.Println("it worked")
	} else {
		fmt.Println("oh no, there's a bug!")
	}
}
