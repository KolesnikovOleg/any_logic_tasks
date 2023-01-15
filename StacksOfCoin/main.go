package main

import "fmt"

func Solution(A []int) int {
	leftPointer := 0
	rightPointer := len(A) - 1
	leftCostSum, rightCostSum := 0, 0
	leftMoveCoinsCount, rightMoveCoinsCount := 0, 0

	doLeftMove := func() {
		leftCostSum += leftMoveCoinsCount
		A[leftPointer] -= leftMoveCoinsCount + leftMoveCoinsCount //not optional
		A[leftPointer+1] += leftMoveCoinsCount
		leftPointer += 1
	}

	doRightMove := func() {
		rightCostSum += rightMoveCoinsCount
		A[rightPointer] -= rightMoveCoinsCount + rightMoveCoinsCount //not optional
		A[rightPointer-1] += rightMoveCoinsCount
		rightPointer -= 1
	}

	for leftPointer != rightPointer {
		leftMoveCoinsCount = int((A[leftPointer] - A[leftPointer]%2) / 2)
		rightMoveCoinsCount = int((A[rightPointer] - A[rightPointer]%2) / 2)
		possibleLeftCost := leftCostSum + leftMoveCoinsCount
		possibleRightCost := rightCostSum + rightMoveCoinsCount

		if possibleRightCost > possibleLeftCost {
			doLeftMove()
			continue
		}
		if possibleRightCost < possibleLeftCost {
			doRightMove()
			continue
		}
		// if equal
		//check next
		if (A[leftPointer+1] + leftMoveCoinsCount) >= (A[rightPointer-1])+rightMoveCoinsCount {
			doLeftMove()
			continue
		}

		doRightMove()
		continue

	}

	return A[leftPointer]
}

func main() {
	// Заменить на нужжый массив для проверки
	A := []int{2, 3, 1, 3}
	fmt.Println(Solution(A))
}
