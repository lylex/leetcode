package main

import (
	"fmt"
	"strings"
)

func convert(s string, numRows int) string {
	if numRows < 2 {
		return s
	}

	width := len(s) / (2*numRows - 2) * (numRows - 1)
	left := len(s) % (2*numRows - 2)
	if left != 0 {
		if left <= numRows {
			width++
		} else {
			width += 1 + left - numRows
		}
	}

	arr := make([][]byte, numRows)
	for i := 0; i < numRows; i++ {
		arr[i] = make([]byte, width)
	}
	fmt.Println(arr)
	var j, p int
	for p < len(s) {
		for k := 0; k < numRows && p < len(s); {
			arr[k][j] = s[p]
			p++
			k++
		}
		j++
		for k := numRows - 2; k > 0 && p < len(s); {
			arr[k][j] = s[p]
			j++
			p++
			k--
		}
	}
	builder := strings.Builder{}
	for i := 0; i < numRows; i++ {
		for j := 0; j < width; j++ {
			if arr[i][j] != 0 {
				builder.WriteByte(arr[i][j])
			}
		}
	}
	return builder.String()
}

func main() {
	fmt.Println(convert("PAYPALISHIRING", 4))
}

