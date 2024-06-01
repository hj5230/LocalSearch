package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

const (
	contextLines      = 1    // Showing lines before and after matches
	sampleSize        = 1024 // Number of bytes to read for content detection
	maxPrintableChars = 0.8  // Percentage threshold for printable characters to consider a file as a text file
)

type StringSlice []string

func (s *StringSlice) String() string {
	return strings.Join(*s, ",")
}

func (s *StringSlice) Set(value string) error {
	*s = append(*s, value)
	return nil
}

func main() {
	helpFlag := flag.Bool("h", false, "Show help information")
	var ignoreDirs StringSlice
	flag.Var(&ignoreDirs, "i", "Directories to ignore (can be specified multiple times)")
	searchStr := flag.String("s", "", "String to search for")
	searchDir := flag.String("d", ".", "Directory to search in")

	flag.Parse()

	if *helpFlag {
		file, err := os.Open("help")
		if err != nil {
			panic(err)
		}

		defer file.Close()

		content, err := io.ReadAll(file)
		if err != nil {
			panic(err)
		}
		fmt.Printf("\n%s\n", string(content))

		os.Exit(0)
	}

	if *searchStr == "" {
		fmt.Println("Specify a string to search for, or use flag -h for help.")
		panic(1)
	}

	err := filepath.Walk(*searchDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		for _, ignoreDir := range ignoreDirs {
			if strings.Contains(path, ignoreDir) {
				if info.IsDir() {
					return filepath.SkipDir
				}
				return nil
			}
		}

		// Temporarily removed the check for text files
		// if !info.IsDir() && isTextFile(path) {
		// 	searchFile(path, *searchStr)
		// }
		if !info.IsDir() {
			searchFile(path, *searchStr)
		}
		return nil
	})

	if err != nil {
		panic(err)
	}
}
