package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

const (
	contextLines      = 1
	sampleSize        = 1024 // Number of bytes to read for content detection
	maxPrintableChars = 0.8  // Percentage threshold for printable characters to consider a file as a text file
)

func main() {
	searchStr := flag.String("s", "", "搜索的字符串")
	searchDir := flag.String("d", ".", "搜索的目录")
	flag.Parse()

	if *searchStr == "" {
		fmt.Println("请指定搜索的字符串")
		os.Exit(1)
	}

	err := filepath.Walk(*searchDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && isTextFile(path) {
			searchFile(path, *searchStr)
		}
		return nil
	})

	if err != nil {
		fmt.Printf("扫描目录出错: %v\n", err)
		os.Exit(1)
	}
}
