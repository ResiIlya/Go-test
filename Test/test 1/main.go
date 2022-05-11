package main

import "fmt"

func main() {
    var a int
	fmt.Scan(&a)

    h := a / 30
    m := a % 30 * 2

	fmt.Println("It is", h, "hours", m, "minutes.")

}