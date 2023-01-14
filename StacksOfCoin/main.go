package main

import "fmt"

func Solution(A []int) int {
	leftPointer := 0
	rightPointer := len(A) - 1
	leftCoastSum, rightCoastSum := 0, 0

	for leftPointer != rightPointer {
		leftMoveCoinsCount := int((A[leftPointer] - A[leftPointer]%2) / 2)
		rightMoveCoinsCount := int((A[rightPointer] - A[rightPointer]%2) / 2)

		if rightMoveCoinsCount > leftMoveCoinsCount {
			leftCoastSum += leftMoveCoinsCount
			A[leftPointer] -= leftMoveCoinsCount + leftMoveCoinsCount
			A[leftPointer+1] += leftMoveCoinsCount
			leftPointer += 1
			continue
		}
		if rightMoveCoinsCount < leftMoveCoinsCount {
			rightCoastSum += rightMoveCoinsCount
			A[rightPointer] -= rightMoveCoinsCount + rightMoveCoinsCount
			A[rightPointer-1] += rightMoveCoinsCount
			rightPointer -= 1
			continue
		}
		// if equal
		//check next
		if (A[leftPointer+1] + leftMoveCoinsCount) >= (A[rightPointer-1])+rightMoveCoinsCount {
			leftCoastSum += leftMoveCoinsCount
			A[leftPointer] -= leftMoveCoinsCount + leftMoveCoinsCount
			A[leftPointer+1] += leftMoveCoinsCount
			leftPointer += 1
			continue
		}

		rightCoastSum += rightMoveCoinsCount
		A[rightPointer] -= rightMoveCoinsCount + rightMoveCoinsCount
		A[rightPointer-1] += rightMoveCoinsCount
		rightPointer -= 1
		continue

	}

	return A[leftPointer]
}

func main() {
	// Заменить на нужжый массив для проверки
	A := []int{1, 1, 0, 1, 1}
	fmt.Println(Solution(A))
}
