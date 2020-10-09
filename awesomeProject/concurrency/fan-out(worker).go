package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(tasksCh <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		task, ok := <-tasksCh
		if !ok {
			return
		}
		//这个sleep这个间隔是可以保护同步的?为什么?
		d := time.Duration(task) * time.Millisecond
		fmt.Println(d)
		//time.Sleep(d) 这个间隔应该是确保是从第一个传递过去，而不是乱传?如果在先传给1的情况下传给了10，如果1还没到10，他就会阻塞掉，直到10先到终点？
		//还是说是模拟完成任务的时间，但这个确实保证了同步喔，换其他就会乱序，但乱序没什么所谓？
		fmt.Println("processing task", task)
	}
}

func pool(wg *sync.WaitGroup, workers, tasks int) {
	tasksCh := make(chan int)

	for i := 0; i < workers; i++ {
		go worker(tasksCh, wg)
	}

	for i := 0; i < tasks; i++ {
		tasksCh <- i
	}

	close(tasksCh)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(36)
	go pool(&wg, 36, 50)
	wg.Wait()
}