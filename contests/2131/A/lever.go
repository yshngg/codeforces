// https://codeforces.com/contest/2131/problem/A

package main

import "fmt"

func main() {
	testCaseNumber := 0
	fmt.Scanln(&testCaseNumber)

	for range testCaseNumber {
		arrayLength := 0
		fmt.Scanln(&arrayLength)

		array := make([]any, arrayLength)
		for i := range array {
			array[i] = new(int)
		}

		fmt.Scanln(array...)
		arrayA := assertArray(array)
		fmt.Scanln(array...)
		arrayB := assertArray(array)

		fmt.Println(lever(arrayA, arrayB))
	}
}

func assertArray(array []any) []int {
	ret := make([]int, len(array))
	for i, v := range array {
		ret[i] = *(v.(*int))
	}
	return ret
}

func lever(arrayA, arrayB []int) int {
	count := 0
	for i := range arrayA {
		if arrayA[i] <= arrayB[i] {
			continue
		}
		count += arrayA[i] - arrayB[i]
	}

	return count + 1
}
