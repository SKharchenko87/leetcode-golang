package p2075

import "strings"

func decodeCiphertext(encodedText string, rows int) string {
	n := len(encodedText)
	cols := n / rows
	decodedBytes := make([]byte, n)
	index := 0
	for i := 0; i < cols; i++ {
		for j := 0; j < rows && i+j < cols; j++ {
			decodedBytes[index] = encodedText[j*cols+j+i]
			index++
		}
	}
	decodedBytes = decodedBytes[:index]
	return strings.TrimRight(string(decodedBytes), " ")
}
