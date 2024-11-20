package costestimate

import (
	"fmt"
)

const pricePerSqFt = 3.80

func RunCostEstimator() {
	var length, width float64

	fmt.Println("Flooring Cost Estimator")
	fmt.Println("Enter the length of the room (in feet): ")
	fmt.Scan(&length)

	fmt.Println("Enter the width of the room (in feet): ")
	fmt.Scan(&width)

	area := length * width
	cost := area * pricePerSqFt

	fmt.Printf("The total cost for flooring a %.2f sq ft area is %.2f\n", area, cost)
}