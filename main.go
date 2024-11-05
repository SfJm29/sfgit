package main

import (
	"fmt"
	"log"
)

func main() {
	var data string
	fmt.Print("Введите целое число: ")
	_, err := fmt.Scanln(&data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Вы ввели число: %v\n", data)
}
