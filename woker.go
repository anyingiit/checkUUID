package main

func createWorker(work func(number int)) chan<- int {
	c := make(chan int)
	go func() {
		for {
			work(<-c)
		}
	}()

	return c
}
