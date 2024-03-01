package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
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
	iterationNumber, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Please provide a valid number")
		return
	}
	for i := 0; i < iterationNumber; i++ {
		x := rand.Float64()
		y := rand.Float64()
		if isDotInCircle(x, y) {
			inCircleDots++
		} else {
			outCircleDots++
		}
	}

	pi := 4 * float64(inCircleDots) / float64(inCircleDots+outCircleDots)
	fmt.Println(pi)
}
