package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func doTask(workerID int, taskName string) {
	fmt.Printf("worker %d is doing task %s\n", workerID, taskName)
	taskDuration := time.Duration(rand.Intn(5))
	time.Sleep(taskDuration) // represents work workerID doing taskName
}

func notifyAllOrWaitForOthers(workerID int, sharedCounter *int, cond *sync.Cond) {
	cond.L.Lock()
	defer cond.L.Unlock()

	// Worker has finished his job so decrements the shared counter
	*sharedCounter--

	// You are the last to complete the task, best notify the rest of the workers.
	if *sharedCounter == 0 {
		fmt.Printf("worker %d is the last to complete work\n", workerID)
		fmt.Printf("last worker signals all workers that they can now proceed with next task\n")
		cond.Broadcast()
		return
	}

	// You finished early, let's wait till the rest are done.
	/*
	   Wait is a special function. Three things will happen upon calling it
	   1. cond.L.Unlock will be called
	   2. scheduler will switch to a different routine and will run it
	   3. Before the wait call returns, cond.L.Lock will be called.
	*/
	cond.Wait()
}

func mainWorker(workerID int, sharedCounter *int, cond *sync.Cond, wg *sync.WaitGroup) {
	// All workers start on task A
	doTask(workerID, "A")
	// Once task A is done, you either wait for others or if your the last
	// to complete the task notify the rest they can start working on task B
	notifyAllOrWaitForOthers(workerID, sharedCounter, cond)
	// All workers start on task B
	doTask(workerID, "B")
	wg.Done()
}

func main() {
	var (
		mu            = &sync.Mutex{}
		cond          = sync.NewCond(mu)
		wg            = &sync.WaitGroup{}
		numWorkers    = 3
		sharedCounter = &numWorkers
	)

	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go mainWorker(i, sharedCounter, cond, wg)
	}
	wg.Wait()
}
