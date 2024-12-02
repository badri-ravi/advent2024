package Day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CheckSorted(arr []int) bool {
	if len(arr) <= 1 {
		return true // A single-element or empty list is considered sorted in ascending
	}

	ascending, descending := true, true

	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			ascending = false
		}
		if arr[i] > arr[i-1] {
			descending = false
		}
	}

	return ascending || descending
}

func absInt(number int) int {
	if number < 0 {
		return -number
	}
	return number
}

func isLevelSafe(num []int) bool {
	flagIncrease, flagDecrease := false, false

	for i := 1; i < len(num); i++ {
		diff := num[i] - num[i-1]

		if diff > 0 {
			flagIncrease = true
		} else if diff < 0 {
			flagDecrease = true
		} else {
			return false
		}

		if flagDecrease && flagIncrease {
			return false
		}

		if diff > 3 || diff < -3 {
			return false
		}
	}

	return true
}

func checkLevelSafeDeletion(num []int) bool {
	for i := 0; i < len(num); i++ {
		isSafe := isLevelSafeDeletion(num, i)
		if isSafe {
			return true
		}
	}
	return false
}

func isLevelSafeDeletion(num []int, i int) bool {
	copyReport := make([]int, len(num))
	copy(copyReport, num)

	if i == len(copyReport)-1 {
		copyReport = copyReport[:i]
	} else {
		copyReport = append(copyReport[:i], copyReport[i+1:]...)
	}
	return isLevelSafe(copyReport)
}

func Day2() {
	readFile, err := os.Open("./Day2/input")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	nums := [][]int{}
	for fileScanner.Scan() {
		tmpNums := []int{}
		strNums := strings.Split(fileScanner.Text(), " ")
		for _, strNum := range strNums {
			i, _ := strconv.Atoi(strNum)
			tmpNums = append(tmpNums, i)
		}
		nums = append(nums, tmpNums)
	}
	safe := 0
	safeWithDeletion := 0
	for _, num := range nums {
		if isLevelSafe(num) {
			safe++
		} else if checkLevelSafeDeletion(num) {
			safeWithDeletion++
		}
	}
	fmt.Println(safe, safe+safeWithDeletion)
}
