package main

import "fmt"

func main() {
	fmt.Println("Welcome to GoLang Experiments!")
	fmt.Println("Enter your name:")
	var name string
	fmt.Scan(&name)
	fmt.Println("Hello", name)
	var age int
	fmt.Println("Enter your age:")
	fmt.Scan(&age)

	if age >= 18 {
		fmt.Println("You are not a kid")
	} else {
		fmt.Println("You are a kid")
	}
	var ans int = check(21, 3)
	fmt.Printf("Answer is: %v yeah", ans)

}

func check(val int, val2 int) int {
	if val > val2 {
		fmt.Println("Val is greater")
	} else {
		fmt.Println("Val2 is greater")
	}
	return (val + val2)
}
