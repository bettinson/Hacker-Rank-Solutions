package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() { //Enter your code here. Read input from STDIN. Print output to STDOUT
	reader := bufio.NewReader(os.Stdin)
	nVal, _ := reader.ReadString('\n')
	arrVals, _ := reader.ReadString('\n')
	arrVals = strings.TrimSuffix(arrVals, "\n")
	result := strings.Split(arrVals, " ")

	var sum int64
	for i := range result {
		num := result[i]
		temp, err := strconv.Atoi(num)
		if err != nil {
			fmt.Print(nVal)
			log.Fatal(err)
		}
		sum = int64(temp) + sum
	}
	fmt.Print(sum)
}
