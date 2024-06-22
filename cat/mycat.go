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

// 	r := strings.NewReader("Тестовое содержимое файла")

// 	var buf bytes.Buffer

// 	if err := readText(r, &buf); err != nil {
// 		fmt.Printf("readText: %s\n", err.Error())
// 	}

// 	fmt.Println(buf.String() == "Тестовое содержимое файла")
// }

func checkArgs(args []string) bool {
	return len(args) == 2
}

func readText(r io.Reader, w io.Writer) error {
	if r == nil {
		return errors.New("reader is nil")
	}
	if w == nil {
		return errors.New("writer is nil")
	}
	_, err := io.Copy(w, r)
	if err != nil {
		fmt.Printf("io.Copy: %s\n", err.Error())
	}

	return nil
}
