package fileprocessor

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

func OpenFileAndReadAllData(fileName string) []byte {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error: ", err)
		return nil
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error: ", err)
		return nil
	}
	return data
}

func ReadPipe() []byte {
	reader := bufio.NewReader(os.Stdin)
	data, err := io.ReadAll(reader)
	if err != nil {
		fmt.Println("Error: ", err)
		return nil
	}
	return data

}

func CountBytes(fileContent []byte) int {
	return len(fileContent)
}

func CountLines(fileContent []byte) int {
	var count int
	reader := bytes.NewReader(fileContent)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		count++
	}
	return count

}

func CountWords(fileContent []byte) int {

	var count int
	reader := bytes.NewReader(fileContent)
	scanner := bufio.NewScanner(reader)
	// change from bufio.ScanLines to bufio.ScanWords as the Scan function
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		count++
	}
	return count

}
