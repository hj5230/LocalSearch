package main

import (
	"os"
	"unicode"
	"unicode/utf8"
)

// Contains functions for text file detection.

func isTextFile(filePath string) bool {
	file, err := os.Open(filePath)
	if err != nil {
		return false
	}
	defer file.Close()

	buffer := make([]byte, sampleSize)
	n, err := file.Read(buffer)
	if err != nil && err.Error() != "EOF" {
		return false
	}

	buffer = buffer[:n]
	return isLikelyText(buffer)
}

func isLikelyText(buffer []byte) bool {
	printableChars := 0
	for i := 0; i < len(buffer); {
		r, size := utf8.DecodeRune(buffer[i:])
		if r == utf8.RuneError && size == 1 {
			return false // Invalid UTF-8 encoding
		}
		if isPrintable(r) {
			printableChars++
		}
		i += size
	}

	return float64(printableChars)/float64(len(buffer)) >= maxPrintableChars
}

func isPrintable(r rune) bool {
	return unicode.IsPrint(r) || unicode.IsSpace(r)
}
