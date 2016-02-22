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
	reader := bufio.NewReader(os.Stdin)
	nVal, _ := reader.ReadString('\n')
	nVal = strings.TrimSuffix(nVal, "\n")
	nValInt, err := strconv.Atoi(nVal)

	if err != nil {
		log.Fatal("Invalid int")
	}

	arr := make([][]int, nValInt)

	// Assuming that the array is a square and gets array from user input
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

	// Calculates diagonal difference
	var diagPos int

}
