package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Sum(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func GetDistance(left []int, right []int) []int {
	slices.Sort(left)
	slices.Sort(right)
	if len(left) != len(right) {
		fmt.Println("Lengths are not equal")
		return nil
	}

	distance := make([]int, len(left))

	for i := 0; i < len(left); i++ {
		distance[i] = left[i] - right[i]
		if distance[i] < 0 {
			distance[i] = -distance[i]
		}
	}
	return distance
}

func GetNumList(fakeNumList []string) []int {
	var numList []int
	for _, list := range fakeNumList {
		intNum, _ := strconv.Atoi(list)
		numList = append(numList, intNum)
	}
	slices.Sort(numList)
	return numList
}

func main() {
	input, err := os.ReadFile("input.txt") // Read the file
	if err != nil {
		fmt.Print(err)
		return // Exit on error
	}

	str := string(input) // Convert content to a 'string'

	// Split the string by new line
	var list = strings.Split(strings.TrimSpace(str), "\n")

	// Clean up spaces in each line
	for i := range list {
		list[i] = strings.ReplaceAll(list[i], "   ", " ") // Replace triple spaces with single space
	}

	var leftList []string
	var rightList []string

	// Parse each line into left and right parts
	for _, line := range list {
		register := strings.Fields(line) // Split by spaces, ignoring extra spaces
		if len(register) >= 2 {          // Ensure there are at least 2 parts
			leftList = append(leftList, register[0])
			rightList = append(rightList, register[1])
		} else {
			fmt.Printf("Skipping invalid line: %q\n", line)
		}
	}

	var leftTiny []int = GetNumList(leftList)
	var rightTiny []int = GetNumList(rightList)

	fmt.Print(Sum(GetDistance(leftTiny, rightTiny)))

}
