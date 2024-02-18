package main

import "fmt"

func main() {
	inputs := []struct {
		numberOfSquares []int32
		birthDay        int32
		birthMonth      int32
		expectedOutput  int32
	}{
		{numberOfSquares: []int32{1, 2, 1, 3, 2}, birthDay: 3, birthMonth: 2, expectedOutput: 2},
		{numberOfSquares: []int32{1, 1, 1, 1, 1, 1}, birthDay: 3, birthMonth: 2, expectedOutput: 0},
		{numberOfSquares: []int32{4}, birthDay: 4, birthMonth: 1, expectedOutput: 1},
		{numberOfSquares: []int32{1, 3, 2, 2, 2, 1}, birthDay: 4, birthMonth: 3, expectedOutput: 0},
		{numberOfSquares: []int32{1, 3, 0, 1, 2, 1, 1, 5, 0, 1}, birthDay: 4, birthMonth: 3, expectedOutput: 4},
		{numberOfSquares: []int32{5, 5, 3, 2, 2, 2, 1, 2, 5, 3, 5, 5, 4, 3, 3, 5}, birthDay: 13, birthMonth: 3, expectedOutput: 3},
	}

	for index, input := range inputs {
		testNumber := index + 1
		output := birthday(input.numberOfSquares, input.birthDay, input.birthMonth)
		if input.expectedOutput != output {
			panic(fmt.Sprintf("Test number %d want %d but got %d", testNumber, input.expectedOutput, output))
		}
		fmt.Printf("Test number %d: success\n", testNumber)
	}
}

func birthday(s []int32, d int32, m int32) int32 {
	var squares int32

	for i := int32(0); i < int32(len(s))-m+1; i++ {
		var count int32

		for j := int32(0); j < m; j++ {
			count += s[i+j]
		}

		if count == d {
			squares++
		}
	}

	return squares
}
