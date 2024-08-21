package main

import (
	"fmt"
	"github.com/mateussssz/multy/processor"
)

func main() {
	fmt.Println("Program has started")
	defer fmt.Println("Program has finished")
	var coresNum, threadsNum int
	fmt.Print("Enter the number of cores: ")
	fmt.Scan(&coresNum)
	fmt.Print("\nEnter the number of threads: ")
	fmt.Scan(&threadsNum)
	cores := make([]int, coresNum) //Создание массива cores
	for idx := range cores {
		cores[idx] = threadsNum
	}
	processor.Processor(cores, coresNum, threadsNum)

}