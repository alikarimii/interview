package main

import "sync"

var counter int

func increment(wg *sync.WaitGroup) {
	for i := 0; i < 1_000_000_000; i++ {
		counter++
	}
	wg.Done()
}
func main2() {
	wg := new(sync.WaitGroup)
	wg.Add(2)
	go increment(wg)
	go increment(wg)
	wg.Wait()
	// what is a counter value here and why? how improve them
}
