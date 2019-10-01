package main

import (
	"fmt"
	"math"
)

func circleArea(radius float64) (float64, error) {
	if radius < 0 {
		return 0, fmt.Errorf("Area calculation failed, radius %0.2f is less than zero", radius)

		//return 0, errors.New("Radius less than zero")
	}

	return math.Pi * radius * radius, nil
}

func main() {
	area, err := circleArea(-2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Area :", area)
}
