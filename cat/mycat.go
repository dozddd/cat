package main

import (
	// "bufio"
	// "bytes"
	"errors"
	"fmt"
	"io"
	"os"
	// "strings"
)

var (
	ReaderIsNilErr = errors.New("reader is nil")
	WriterIsNilErr = errors.New("writer is nil")
)

func main() {
	if !checkArgs(os.Args) {
		fmt.Println("Укажите только один файл")
		return
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("os.Open: %s", err.Error())
		return
	}
	defer file.Close()

	if err := readText(file, os.Stdout); err != nil {
		fmt.Printf("readText: %s\n", err.Error())
	}
}

func checkArgs(args []string) bool {
	return len(args) == 2
}

func readText(r io.Reader, w io.Writer) error {
	if r == nil {
		return ReaderIsNilErr
	}

	if w == nil {
		return WriterIsNilErr
	}

	_, err := io.Copy(w, r)
	if err != nil {
		return fmt.Errorf("io.Copy: %w", err)
	}

	return nil
}
