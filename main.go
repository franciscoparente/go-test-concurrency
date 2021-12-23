package main

import (
	"fmt"
	"strconv"
	"sync"
	"sync/atomic"
)

func main() {
	asyncSolution()
	asyncProblem()
	concurrencyProblem()
	concurrencySolution()
}

func asyncSolution() {
	var wg sync.WaitGroup
	for i := 0; i <= 5; i++ {
		fmt.Println("before")
		wg.Add(1)
		go func() {
			for i := 0; i <= 100000; i++ {
				fmt.Println("after " + strconv.Itoa(i))
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func asyncProblem() {
	for i := 0; i <= 5; i++ {

		fmt.Println("before")
		go func() {
			for i := 0; i <= 100000; i++ {
				fmt.Println("after " + strconv.Itoa(i))
			}
		}()
	}
}

func concurrencyProblem() {
	var wg sync.WaitGroup

	j := uint64(0)
	for i := 0; i <= 10000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			j = j + 1
		}()
	}

	wg.Wait()
	fmt.Println("result ", j)

}

func concurrencySolution() {
	var wg sync.WaitGroup

	j := uint64(0)
	for i := 0; i <= 10000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddUint64(&j, 1)
		}()
	}

	wg.Wait()
	fmt.Println("result ", j)
}
