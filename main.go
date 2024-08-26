package main

import (
	"fmt"

	"github.com/mateussssz/multy/processor"
	st "github.com/mateussssz/multy/structs"
)

func main() {
	var coresNum, threadsNum = start_program()
	defer finish_program()

	cLoad := &st.CoresLoad{Cores: make([]int, coresNum)}
	for idx := range cLoad.Cores {
		cLoad.Cores[idx] = threadsNum
	}

	processor.Processor(cLoad, coresNum, threadsNum)

}

func start_program() (int, int) {
	var coresNum, threadsNum int
	fmt.Println("Program has started")
	fmt.Print("Enter the number of cores: ")
	fmt.Scan(&coresNum)
	fmt.Print("\nEnter the number of threads: ")
	fmt.Scan(&threadsNum)
	return coresNum,threadsNum
}

func finish_program(){
	fmt.Println("Program has finished")
}