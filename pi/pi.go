package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

var (
	inCircleDots  = 0
	outCircleDots = 0
)

func calculateDistance(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}

func isDotInCircle(x, y float64) bool {
	return calculateDistance(x, y) <= 1
}

func main() {
	var rw sync.RWMutex
	iterationNumber, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Please provide a valid number")
		return
	}
	startTime := time.Now()
	var wg sync.WaitGroup
	for i := 0; i < 8; i++ {
		wg.Add(1)
		go func() {
			var inCircleDotss, outCircleDotss int
			for i := 0; i < iterationNumber/10; i++ {
				x := rand.Float64()
				y := rand.Float64()
				if isDotInCircle(x, y) {
					inCircleDotss++
				} else {
					outCircleDotss++
				}
			}
			rw.Lock()
			inCircleDots += inCircleDotss
			outCircleDots += outCircleDotss
			rw.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Time taken: ", time.Since(startTime))

	pi := 4 * float64(inCircleDots) / float64(inCircleDots+outCircleDots)
	fmt.Println(pi)
}
