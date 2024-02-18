package main

import "fmt"

func main() {
	// output := birthday([]int32{1, 2, 1, 3, 2}, 3, 2) // expected: 2
	// output := birthday([]int32{1, 1, 1, 1, 1, 1}, 3, 2) // expected: 0
	// output := birthday([]int32{4}, 4, 1) // expected: 1
	// output := birthday([]int32{1, 3, 2, 2, 2, 1}, 4, 3)    // expected: 3
	// output := birthday([]int32{1, 3, 2, 2, 2, 1, 3}, 4, 3) // expected: 4
	output := birthday([]int32{5, 5, 3, 2, 2, 2, 1, 2, 5, 3, 5, 5, 4, 3, 3, 5}, 13, 3) // expected: 3
	fmt.Println(output)
}

func birthday(s []int32, d int32, m int32) int32 {
	squares := int32(0)
	lenght := len(s)

	if lenght == 1 && s[0] == d {
		return 1
	}

	for i := 0; i < lenght; i++ {
		if i+1 == lenght {
			break
		}

		if s[i]+s[i+1] == d {
			squares++
		}
	}

	return squares
}
