package main

import "fmt"

//User its type.
type User struct {
	id   int64
	name string
}

func main() {
	fmt.Println("hello world")

	user := User{
		12,
		"Alex",
	}

	user.id = 10

	fmt.Println(user)
}

//Inc its summ a variable plus b variable.
func Inc(a, b int) int {
	return a + b
}
