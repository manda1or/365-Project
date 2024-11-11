package reviews

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ReviewStruct struct {
	ReviewerName  string
	ReviewComment string
	ReviewRating  int
}

func CreateNewReview() ReviewStruct {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Make Your Review!")

	// Reviewer Name Input
	fmt.Println("Enter Your Name:")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	//Reviewer Comment Input
	fmt.Println("Enter Your Comment:")
	comment, _ := reader.ReadString('\n')
	comment = strings.TrimSpace(comment)

	// Reviewer Rating
	var rating int
	for {
		fmt.Println("Enter a rating 1-5:")
		ratingInput, _ := reader.ReadString('\n')
		ratingInput = strings.TrimSpace(ratingInput)

		//Turn string input into int
		var err error
		rating, err = strconv.Atoi(ratingInput)

		if err != nil || rating < 1 || rating > 5 {
			fmt.Println("Invalid input. Enter a valid rating (1-5).")
			continue
		}
		break
	}

	return ReviewStruct{
		ReviewerName:  name,
		ReviewComment: comment,
		ReviewRating:  rating,
	}
}
