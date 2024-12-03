package Day3

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func Day3() {
	content, err := os.ReadFile("./Day3/input")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	memory := string(content)

	mulPattern := regexp.MustCompile(`mul\(([0-9]{1,3}),([0-9]{1,3})\)|don't\(\)|do\(\)`)
	isMulEnabled := true
	total := 0
	matches := mulPattern.FindAllStringSubmatch(memory, -1)
	for _, match := range matches {
		switch match[0] {
		case "do()":
			isMulEnabled = true
		case "don't()":
			isMulEnabled = false
		default:
			if isMulEnabled {
				num1, _ := strconv.Atoi(match[1])
				num2, _ := strconv.Atoi(match[2])
				total += num1 * num2
			}
		}
	}
	fmt.Println(total)
}
