package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

// ReadTabSeparatedFile reads a tab-separated file and stores the data in a 2D slice
func ReadTabSeparatedFile(filePath string) ([][]int, error) {
	var data [][]int

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Split the line into fields by tab
		line := scanner.Text()
		fields := strings.Split(line, "   ")
		var row []int
		for _, field := range fields {
			value, err := strconv.Atoi(field)
			if err != nil {
				return nil, fmt.Errorf("failed to convert value '%s' to int: %v", field, err)
			}
			row = append(row, value)
		}
		data = append(data, row)
	}

	// Check for scanning errors
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	return data, nil
}

func main() {

	input, err := ReadTabSeparatedFile("day1/input.txt")
	if err == nil {
		// getDifference(input)
		start := time.Now()
		getSimilarityScore(input)
		elapsed := time.Since(start)
		fmt.Printf("\nElapsed time %s", elapsed)
	}else{
		fmt.Print(err)
	}
	}


 func getDifference(input [][]int){

	firstCol := make([]int, 0)
	secondCol := make([]int, 0)
	for rowIdx := 0; rowIdx<len(input); rowIdx++{
		val := input[rowIdx][0]
		firstCol = append(firstCol,val)
		secondCol = append(secondCol, input[rowIdx][1])
	}

	sort.Ints(firstCol)
	sort.Ints(secondCol)

	total := 0
	for rowIdx := 0; rowIdx<len(firstCol); rowIdx++{
		diff := 0
		if secondCol[rowIdx] > firstCol[rowIdx] {
			diff = secondCol[rowIdx] - firstCol[rowIdx]
		}else{
			diff = firstCol[rowIdx] - secondCol[rowIdx]
		}
		total += diff
	}
	fmt.Print(total)

}
func remove(s []int, i int) []int {
    s[i] = s[len(s)-1]
    return s[:len(s)-1]
}

func countNumOccurences(col *[]int, number int)(int){
	count := 0
	slice := *col
	idxToRemove := make(map[int]bool)
	for idx := 0; idx < len(slice); idx++{
		if slice[idx] == number{
			count++
			idxToRemove[idx] = true
		}
	}
	for idx :=0;idx<len(idxToRemove);idx++{
		if idxToRemove[idx]{
			slice[idx] = slice[len(slice)-1] 
			slice = slice[:len(slice)-1]
		}
	}
	*col = slice
	return count
}

func getSimilarityScore(input [][]int){

	firstCol := make([]int, 0)
	secondCol := make([]int, 0)
	for rowIdx := 0; rowIdx<len(input); rowIdx++{
		val := input[rowIdx][0]
		firstCol = append(firstCol,val)
		secondCol = append(secondCol, input[rowIdx][1])
	}

	numberMap := make(map[int]int)

	rowIdx := 0
	total := 0
	for rowIdx <len(firstCol){
		_,valInMap := numberMap[firstCol[rowIdx]]
		if !valInMap{
			numberMap[firstCol[rowIdx]] = countNumOccurences(&secondCol, firstCol[rowIdx])
		}
		
		total += firstCol[rowIdx] * numberMap[firstCol[rowIdx]]
		rowIdx++
	}

	print(total)
}