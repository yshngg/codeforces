// https://codeforces.com/contest/2131/problem/B

package main

import "fmt"

func main() {
	testCaseNumber := 0
	fmt.Scanln(&testCaseNumber)

	for range testCaseNumber {
		arrayLength := 0
		fmt.Scan(&arrayLength)

		array := make([]int, arrayLength)
		for i := range arrayLength - 1 {
			if i%2 == 0 {
				array[i] = -1
				continue
			}
			array[i] = 3
		}

		if arrayLength%2 == 0 {
			array[arrayLength-1] = 2
		} else {
			array[arrayLength-1] = -1
		}

		for i := range arrayLength - 1 {
			fmt.Printf("%d ", array[i])
		}
		fmt.Printf("%d\n", array[arrayLength-1])
	}
}
