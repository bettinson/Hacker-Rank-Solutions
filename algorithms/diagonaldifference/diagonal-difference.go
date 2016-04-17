package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	arr, nValInt := getInput()
	diff := diagonalDifference(arr, nValInt)
	fmt.Print(diff)
	// Calculates diagonal difference
}

func diagonalDifference(arr [][]int, nValInt int) int {
	var leftDiagSum int
	var rightDiagSum int
	for i := 0; i < nValInt; i++ {
		leftDiagSum += arr[i][i]
		fmt.Println(arr[i][i])
		rightDiagSum += arr[nValInt-i-1][i]
		fmt.Println(arr[nValInt-i-1][nValInt-i-1])
	}
	fmt.Println(rightDiagSum)
	fmt.Println(leftDiagSum)
	return (leftDiagSum - rightDiagSum)
}

func getInput() ([][]int, int) {
	reader := bufio.NewReader(os.Stdin)
	nVal, _ := reader.ReadString('\n')
	nVal = strings.TrimSuffix(nVal, "\n")
	nValInt, err := strconv.Atoi(nVal)

	if err != nil {
		log.Fatal("Invalid int")
	}

	arr := make([][]int, nValInt)

	// Assuming that the array is a square gets array from user input
	for i := 0; i < nValInt; i++ {
		arrVals, _ := reader.ReadString('\n')
		arrVals = strings.TrimSuffix(arrVals, "\n")
		result := strings.Split(arrVals, " ")

		arrValsInt := make([]int, nValInt)

		for j := range result {
			num := result[j]
			temp, err := strconv.Atoi(num)
			if err != nil {
				fmt.Print(nVal)
				log.Fatal(err)
			}
			arrValsInt[j] = temp
		}
		arr[i] = arrValsInt
	}

	return arr, nValInt
}
