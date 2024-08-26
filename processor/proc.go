package processor

import (
	"fmt"
	"runtime"
	"time"

	"github.com/mateussssz/multy/keyboard"
	st "github.com/mateussssz/multy/structs"
)

func Processor(cLoad *st.CoresLoad, coresNum int, threadsNum int) {
	keyChannel := make(chan uint32, 1)
	var taskChannels []chan uint32 = make([]chan uint32, coresNum)
	for i := 0; i < coresNum; i++ {
		taskChannels[i] = make(chan uint32, threadsNum)
	}
	go keyboard.KeyboardHandler(keyChannel)
	var queue []uint32

	for {
		if len(keyChannel) != 0 {
			PID := <-keyChannel
			if PID == 0 {
				close(keyChannel)
			}
			queue = append(queue, PID)
		}

		tasksLen := len(queue)
		for tasksLen != 0 { //Подумать, как обрабатывать процессы
			PID := queue[0] //И принимать завершенную работу
			if PID == 0 {
				waitForCores()
				return
			}

			freeCore := 0
			freeThreads := cLoad.Cores[0]
			for i := 1; i < len(cLoad.Cores); i++ {
				if cLoad.Cores[i] > freeThreads {
					freeCore = i
					freeThreads = cLoad.Cores[i]
				}
			}
			if cLoad.Cores[freeCore] != 0 {
				if cLoad.Cores[freeCore] == threadsNum {
					go core(freeCore, threadsNum, taskChannels[freeCore], cLoad)
				}
				taskChannels[freeCore] <- PID
				queue = queue[1:]
			} else {
				break
			}
			tasksLen--
		}
	}
}

func core(coreNumber int, threadsNum int, taskChannel chan uint32, cLoad *st.CoresLoad) {
	fmt.Println(coreNumber, "core was activated")
	freeChannel := make(chan int, threadsNum)
	var freeThreads []int
	for i := 0; i < threadsNum; i++ {
		freeThreads = append(freeThreads, i)
	}
	defer fmt.Println(coreNumber, "core was disabled because of job`s lack")
	for {
		for len(taskChannel) > 0 && len(freeThreads) > 0 {
			threadNumber := freeThreads[0]
			PID := <-taskChannel
			freeThreads = freeThreads[1:]
			cLoad.Lock()
			cLoad.Cores[coreNumber]-- //Добавление освободившихся потоков в стек freeThreads и посыл процессору, сколько у нас свободных мест
			cLoad.Unlock()
			go thread(coreNumber, threadNumber, PID, freeChannel)
		}
		for len(freeChannel) != 0 {
			value := <-freeChannel
			freeThreads = append(freeThreads, value)
			cLoad.Lock()
			cLoad.Cores[coreNumber]++ //Добавление освободившихся потоков в стек freeThreads и посыл процессору, сколько у нас свободных мест
			cLoad.Unlock()
		}

		if len(freeThreads) == threadsNum && len(taskChannel) == 0 {
			close(freeChannel)
			break
		} //Если новых задач нет и длина стека равна количеству потоков - break из цикла
	}
}

func thread(coreNumber int, threadNumber int, PID uint32, freeChannel chan int) {
	fmt.Println(coreNumber, "core started the", threadNumber, "thread(the process number", PID, ")")
	for i := 1; i < 5; i++ {
		time.Sleep(time.Second)
		fmt.Println(coreNumber, "core,", threadNumber, "thread:", 5-i, "seconds left")
	}
	fmt.Println(coreNumber, "core finished the", threadNumber, "thread")
	freeChannel <- threadNumber
}

func waitForCores() {
	fmt.Println("You pressed ESC! Waiting for cores to finish their job...")
	defer fmt.Println("All cores finished their jobs. Escape!")
	for runtime.NumGoroutine() != 1 {
		time.Sleep(time.Second)
	}
}
