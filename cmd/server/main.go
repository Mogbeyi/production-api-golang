package main

import "fmt"

func Run() error {
	fmt.Println("Starting up our application...")
	return nil
}

func main() {
	fmt.Println("Go rest API course")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
