package main

import (
	"fmt"
	"github.com/Goku-kun/wc-golang/lib/fileprocessor"
	"github.com/Goku-kun/wc-golang/lib/flagparser"
)

func main() {
	flags, fileNames := flagparser.ParseFlags()

	if len(fileNames) == 0 {
		fileContent := fileprocessor.ReadPipe()
		l := fileprocessor.CountLines(fileContent)
		w := fileprocessor.CountWords(fileContent)
		b := fileprocessor.CountBytes(fileContent)

		if flags["c"] {
			fmt.Printf("    %v\n", b)
		} else if flags["l"] {
			fmt.Printf("    %v\n", l)
		} else if flags["w"] {
			fmt.Printf("    %v\n", w)
		} else {
			fmt.Printf("    %v %v %v \n", l, w, b)
		}
		return
	}

	for _, fileName := range fileNames {
		if flags["c"] {
			b := fileprocessor.CountBytes(fileprocessor.OpenFileAndReadAllData(fileName))
			fmt.Printf("    %v %v\n", b, fileName)
		} else if flags["l"] {
			l := fileprocessor.CountLines(fileprocessor.OpenFileAndReadAllData(fileName))
			fmt.Printf("    %v %v\n", l, fileName)
		} else if flags["w"] {
			w := fileprocessor.CountWords(fileprocessor.OpenFileAndReadAllData(fileName))
			fmt.Printf("    %v %v\n", w, fileName)
		} else {
			b := fileprocessor.CountBytes(fileprocessor.OpenFileAndReadAllData(fileName))
			l := fileprocessor.CountLines(fileprocessor.OpenFileAndReadAllData(fileName))
			w := fileprocessor.CountWords(fileprocessor.OpenFileAndReadAllData(fileName))
			fmt.Printf("    %v %v %v %v\n", l, w, b, fileName)
		}
	}
}
