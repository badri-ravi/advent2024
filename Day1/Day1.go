package Day1

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func absInt(number int) int {
	if number < 0 {
		return -number
	}
	return number
}

func countOccurrences(arr []int, target int) int {
	count := 0
	for _, value := range arr {
		if value == target {
			count++
		}
	}
	return count
}

func Day1part1() {
	readFile, err := os.Open("./Day1/input")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	nums1 := []int{}
	nums2 := []int{}
	for fileScanner.Scan() {
		num1, _ := strconv.Atoi(strings.Split(fileScanner.Text(), "   ")[0])
		num2, _ := strconv.Atoi(strings.Split(fileScanner.Text(), "   ")[1])
		nums1 = append(nums1, num1)
		nums2 = append(nums2, num2)
	}
	sort.Ints(nums1)
	sort.Ints(nums2)
	tot := 0
	for i := 0; i < len(nums1); i++ {
		tot += absInt(nums1[i] - nums2[i])
	}
	fmt.Println(tot)
}

func Day1part2() {
	readFile, err := os.Open("./Day1/input")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	nums1 := []int{}
	nums2 := []int{}
	for fileScanner.Scan() {
		num1, _ := strconv.Atoi(strings.Split(fileScanner.Text(), "   ")[0])
		num2, _ := strconv.Atoi(strings.Split(fileScanner.Text(), "   ")[1])
		nums1 = append(nums1, num1)
		nums2 = append(nums2, num2)
	}
	tot := 0
	for i := 0; i < len(nums1); i++ {

		tot += nums1[i] * countOccurrences(nums2, nums1[i])

	}
	fmt.Println(tot)
}
