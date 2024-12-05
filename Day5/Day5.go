package Day5

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type PageNum struct {
	X int
	Y int
}

func findInPage(pageRules []PageNum, pageX int, pageY int) bool {
	for _, pageRule := range pageRules {
		if pageRule.X == pageX && pageRule.Y == pageY {
			return true
		}
	}
	return false
}

func findInPages(pageRules []PageNum, pageX []int, pageY []int, target int) bool {
	countX := 0
	countY := 0
	for i := 0; i < len(pageX); i++ {
		for _, pageRule := range pageRules {
			if pageRule.X == pageX[i] && pageRule.Y == target {
				countX++
			}
		}
	}
	for i := 0; i < len(pageY); i++ {
		for _, pageRule := range pageRules {
			if pageRule.X == target && pageRule.Y == pageY[i] {
				countY++
			}
		}
	}

	return countX == len(pageX) && countY == len(pageY)
}

func MiddleElement(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	middleIndex := len(arr) / 2
	if len(arr)%2 == 0 {
		middleIndex-- // For even-length arrays, choose the lower middle.
	}
	return arr[middleIndex]
}

func Day5() {
	filePath := "./Day5/input" // Replace with your file path
	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading the file:", err)
		return
	}
	pageRules := []PageNum{}
	sections := strings.Split(string(content), "\n\n")
	for _, rules := range strings.Split(sections[0], "\n") {
		page := strings.Split(rules, "|")
		pageNum1, _ := strconv.Atoi(page[0])
		pageNum2, _ := strconv.Atoi(page[1])
		pageRules = append(pageRules, PageNum{pageNum1, pageNum2})
	}
	pageNumberOrders := [][]int{}
	for _, pageOrders := range strings.Split(sections[1], "\n") {
		order := []int{}
		for _, pageOrder := range strings.Split(pageOrders, ",") {
			num, _ := strconv.Atoi(pageOrder)
			order = append(order, num)
		}
		pageNumberOrders = append(pageNumberOrders, order)
	}
	totalPart1 := 0
	totalPart2 := 0
	for i := 0; i < len(pageNumberOrders); i++ {
		count := 0
		for j := 0; j < len(pageNumberOrders[i]); j++ {
			if findInPages(pageRules, pageNumberOrders[i][:j], pageNumberOrders[i][j+1:], pageNumberOrders[i][j]) {
				count++
			}
		}
		if count == len(pageNumberOrders[i]) {
			totalPart1 += MiddleElement(pageNumberOrders[i])
		}

	}
	fmt.Println(totalPart1)
	for _, update := range pageNumberOrders {
		original := make([]int, len(update))
		copy(original, update)

		slices.SortFunc(update, func(a, b int) int {
			if findInPage(pageRules, a, b) {
				return -1
			} else {
				return 1
			}
		})

		if !slices.Equal(update, original) {
			totalPart2 += update[len(update)/2]
		}
	}
	fmt.Println(totalPart2)
}
