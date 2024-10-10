package main

import "fmt"

func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	resultChan := make(chan int)

	go func() {
		defer close(resultChan)

		for {
			select {
			case firstVal, ok := <-firstChan:
				if !ok {
					return
				}
				resultChan <- firstVal * firstVal
			case secondVal, ok := <-secondChan:
				if !ok {
					return
				}
				resultChan <- secondVal * 3
			case <-stopChan:
				return
			}
		}
	}()

	return resultChan
}

func main() {
	firstChan := make(chan int)
	secondChan := make(chan int)
	stopChan := make(chan struct{})

	go func() {
		firstChan <- 5
		close(firstChan)
	}()

	go func() {
		secondChan <- 10
		close(secondChan)
	}()

	go func() {
		stopChan <- struct{}{}
		close(stopChan)
	}()

	resultChan := calculator(firstChan, secondChan, stopChan)
	for result := range resultChan {
		fmt.Println("Result:", result)
	}
}
