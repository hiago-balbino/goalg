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

	for _, input := range inputs {
		output := birthday(input.numberOfSquares, input.birthDay, input.birthMonth)
		if input.expectedOutput != output {
			panic(fmt.Sprintf("Want %d but got %d", input.expectedOutput, output))
		}
	}
}

func birthday(s []int32, d int32, m int32) int32 {
	if len(s) == 1 && s[0] == d {
		return 1
	}

	var squares int32
	for i := int32(0); i < int32(len(s))-m; i++ {

		var count int32
		for j := i; j < i+m; j++ {
			count += s[j]
		}

		if count == d {
			squares++
		}
	}

	return squares
}
