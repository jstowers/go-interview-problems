package main

import "fmt"

func main() {
	// Expected = 12
	// cardPoints := []int{1, 2, 3, 4, 5, 6, 1}

	// Expected = 248
	//cardPoints := []int{100, 40, 17, 9, 73, 75}
	//k := 3

	// Expected =
	//cardPoints := []int{9, 7, 7, 9, 7, 7, 9}

	// Expected = 202
	//cardPoints := []int{1, 79, 80, 1, 1, 1, 200, 1}
	//k := 3

	// Expected = 232
	cardPoints := []int{11, 49, 100, 20, 86, 29, 72}
	k := 4

	score := maxScore(cardPoints, k)

	fmt.Println("score = ", score)

}

func maxScore(cardPoints []int, k int) int {
	fmt.Println("\ncardPoints[] =", cardPoints)
	highScore := 0
	leftIndex := k

	for leftIndex >= 0 {
		fmt.Println("------------------------------")
		rightIndex := k - leftIndex
		fmt.Println("leftIndex =", leftIndex, "rightIndex =", rightIndex)
		leftSum := 0
		rightSum := 0

		// select and sum values for left side of window
		if leftIndex > 0 {
			leftArray := cardPoints[0:leftIndex]
			fmt.Println("leftArray =", leftArray)
			for i := 0; i < len(leftArray); i++ {
				leftSum += leftArray[i]
			}
		}

		// select and sum values for right side of window
		if rightIndex > 0 {
			rightArray := cardPoints[len(cardPoints)-rightIndex:]
			fmt.Println("rightArray = ", rightArray)
			for i := 0; i < len(rightArray); i++ {
				rightSum += rightArray[i]
			}
		}

		rowSum := leftSum + rightSum
		fmt.Println("rowSum =", rowSum)

		if rowSum > highScore {
			highScore = rowSum
		}

		leftIndex--
	}

	return highScore
}

// func recursive(cardPoints []int, k int, counter int, sum int) int {
// 	fmt.Println("------------------------------")
// 	fmt.Println("counter =", counter)
// 	fmt.Println("sum =", sum)

// 	if counter == k {
// 		fmt.Println("sum final =", sum)
// 		return sum
// 	}

// 	lowIndex := 0
// 	highIndex := len(cardPoints) - 1

// 	// case 1 : lowIndex > highIndex
// 	if cardPoints[lowIndex] > cardPoints[highIndex] {
// 		fmt.Println("INSIDE LOW INDEX")
// 		sum += cardPoints[lowIndex]
// 		cardPoints = cardPoints[lowIndex+1:]
// 		fmt.Println("cardPoints =", cardPoints)
// 		return recursive(cardPoints, k, counter+1, sum)

// 	} else if cardPoints[highIndex] > cardPoints[lowIndex] {
// 		fmt.Println("INSIDE HIGH INDEX")
// 		fmt.Println("cardPoints[] =", cardPoints)
// 		fmt.Println("highIndex =", highIndex)
// 		sum += cardPoints[highIndex]
// 		cardPoints = cardPoints[0:highIndex]
// 		fmt.Println("cardPoints[] after =", cardPoints)
// 		return recursive(cardPoints, k, counter+1, sum)

// 	} else {
// 		fmt.Println("INSIDE EQUAL")
// 		sumTemp := cardPoints[lowIndex]
// 		fmt.Println("sumTemp =", sumTemp)
// 		cardPointsLow := cardPoints[lowIndex:highIndex]
// 		fmt.Println("cardPointsLow =", cardPointsLow)
// 		cardPointsHigh := cardPoints[lowIndex+1:]
// 		fmt.Println("cardPointsHigh =", cardPointsHigh)

// 		sumLow := recursive(cardPointsLow, k, counter+1, sumTemp)
// 		sumHigh := recursive(cardPointsHigh, k, counter+1, sumTemp)

// 		fmt.Println("sumLow =", sumLow)
// 		fmt.Println("sumHigh =", sumHigh)

// 		if sumLow >= sumHigh {
// 			sum += sumLow
// 		} else {
// 			sum += sumHigh
// 		}
// 	}
// 	return sum
// }

// func maxScore(cardPoints []int, k int) int {
// 	return recursive(cardPoints, k, 0, 0)
// }

/*
func maxScore(cardPoints []int, k int) int {

	sum := 0
	counter := 0

	lowIndex := 0
	highIndex := len(cardPoints) - 1

	for counter < k {
		if cardPoints[lowIndex] == cardPoints[highIndex] {
			fmt.Println("INSIDE EQUAL INDEX")
			sum += cardPoints[lowIndex]

			// increase lowIndex by 1 and compare this element to highIndex
			maxScore(cardPoints[lowIndex + 1, highIndex], k-1)

			// decrease highIndex by 1 and compare to lowIndex
			maxScore(cardPoints[lowIndex, highIndex - 1], k-1)

			lowIndex += 1
			highIndex -= 1
		} else if cardPoints[lowIndex] > cardPoints[highIndex] {
			fmt.Println("INSIDE LOW INDEX")
			fmt.Println("lowIndex = ", lowIndex, "highIndex =", highIndex)
			fmt.Println("cardPoints[lowIndex] = ", cardPoints[lowIndex], "cardPoints[highIndex] = ", cardPoints[highIndex])
			sum += cardPoints[lowIndex]
			lowIndex += 1
		} else {
			fmt.Println("INSIDE HIGH INDEX")
			sum += cardPoints[highIndex]
			highIndex -= 1
		}
		counter += 1
	}
	return sum
}
*/
