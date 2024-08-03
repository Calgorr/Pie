package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	var n int64
	fmt.Print("Please input the number of iterations: ")
	fmt.Scan(&n)

	step := 1.0 / float64(n)
	numCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPU)
	var wg sync.WaitGroup
	sumChannel := make(chan float64, numCPU)

	start := time.Now()

	for i := 0; i < numCPU; i++ {
		wg.Add(1)
		go func(start, end int64) {
			defer wg.Done()
			localSum := 0.0
			for j := start; j < end; j++ {
				x := (float64(j) + 0.5) * step
				localSum += 4.0 / (1.0 + x*x)
			}
			sumChannel <- localSum
		}(int64(i)*n/int64(numCPU), int64(i+1)*n/int64(numCPU))
	}

	go func() {
		wg.Wait()
		close(sumChannel)
	}()

	sum := 0.0
	for localSum := range sumChannel {
		sum += localSum
	}

	pi := step * sum
	elapsed := time.Since(start)

	fmt.Printf("Pi is approximately %.15f\n", pi)
	fmt.Printf("Time: %s\n", elapsed)
}
