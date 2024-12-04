package Day4

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func checkHorizontal(g [][]string) int {
	var found int = 0
	for i := 0; i < len(g); i++ {
		for j := 0; j < len(g[i])-3; j++ {
			letters := strings.Join(g[i][j:j+4], "")
			if letters == "XMAS" || letters == "SAMX" {
				found += 1
			}
		}
	}
	return found
}

func checkVertical(g [][]string) int {
	var found int = 0
	for col := 0; col < len(g[0]); col++ {
		for startRow := 0; startRow < len(g)-3; startRow++ {
			var letters []string
			for i := 0; i < 4; i++ {
				letters = append(letters, g[startRow+i][col])
			}
			word := strings.Join(letters, "")
			if word == "XMAS" || word == "SAMX" {
				found += 1
			}
		}
	}
	return found
}

func checkDiagonal(g [][]string) int {
	var found int = 0
	//up to the right
	for col := 3; col < len(g[0]); col++ {
		for startRow := 0; startRow < len(g)-3; startRow++ {
			var letters []string
			for i := 0; i < 4; i++ {
				letters = append(letters, g[startRow+i][col-i])
			}
			word := strings.Join(letters, "")
			if word == "XMAS" || word == "SAMX" {
				found += 1
			}
		}
	}
	//up to the left
	for col := 0; col < len(g[0])-3; col++ {
		for startRow := 0; startRow < len(g)-3; startRow++ {
			var letters []string
			for i := 0; i < 4; i++ {
				letters = append(letters, g[startRow+i][col+i])
			}
			word := strings.Join(letters, "")
			if word == "XMAS" || word == "SAMX" {
				found += 1
			}
		}
	}
	return found
}

func findXmasCount(g [][]string) int {
	var found int = 0
	for startCol := 0; startCol <= len(g[0])-3; startCol++ {
		for startRow := 0; startRow <= len(g)-3; startRow++ {
			var goRight []string
			var goLeft []string
			for i := 0; i < 3; i++ {
				goRight = append(goRight, g[startRow+i][startCol+i])
				goLeft = append(goLeft, g[startRow+2-i][startCol+i])
			}
			wordRight := strings.Join(goRight, "")
			wordLeft := strings.Join(goLeft, "")
			if (wordRight == "MAS" || wordRight == "SAM") && (wordLeft == "MAS" || wordLeft == "SAM") {
				found += 1
			}
		}
	}
	return found
}

func Day4() {
	readFile, err := os.Open("./Day4/input")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	words := [][]string{}
	for fileScanner.Scan() {
		line := fileScanner.Text()
		tmpWord := []string{}
		for _, word := range strings.Split(line, "") {
			tmpWord = append(tmpWord, word)
		}
		words = append(words, tmpWord)
	}
	var part1 int = checkHorizontal(words) + checkVertical(words) + checkDiagonal(words)
	part2 := findXmasCount(words)
	fmt.Println(part1, part2)

}
