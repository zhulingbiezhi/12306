package main

import "fmt"

func main() {
	err := loginStart()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
