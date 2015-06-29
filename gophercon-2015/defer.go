package main

func typicalFunction() {
	defer func() {
		fmt.Println("I am deferred")
	}()

	fmt.Println("ret now")
	return
}

func main() {
	typicalFunction()
}
