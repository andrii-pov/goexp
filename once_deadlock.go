package main

import "sync"

func main() {

	var onceA, onceB sync.Once
	var initB func()
	initA := func() { onceB.Do(initB) }
	initB = func() { onceA.Do(initA) }
	// This will cause a deadlock as when onceB.Do is called, we already called onceA.Do but it hasn't returned yet, so we are waiting for onceA.Do to finish
	onceA.Do(initA)

}
