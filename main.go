package main

import (
	"fmt"
	"sync"
)

func workerFunc(wg *sync.WaitGroup, tasksChannel chan int) {
	defer wg.Done()
	for val := range tasksChannel {
		fmt.Printf("worker started %d", val)
	}

}
func main() {
	const workers = 3
	const tasks = 10
	wg := sync.WaitGroup{}
	tasksCh := make(chan int, 10)
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go workerFunc(&wg, tasksCh)
	}
	go func() {
		for i := 1; i <= 10; i++ {
			tasksCh <- i
		}
	}()
	go func() {
		wg.Wait()
		close(tasksCh)
	}()
}
