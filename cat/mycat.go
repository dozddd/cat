package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if checkArgs(os.Args) == false {
		fmt.Println("Укажите только один файл")
	} else {
		err := readText(os.Args[1])
		if err != nil {
			fmt.Println("Проверьте правильность указанного файла")
		}
	}
}

func checkArgs(args []string) bool {
	numberOfArgs := len(args)
	if numberOfArgs == 2 {
		return true
	} else {
		return false
	}
}

func readText(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	fmt.Println("Чтение файла...")
	Data := bufio.NewScanner(file)
	for Data.Scan() {
		fmt.Printf("%s\n", Data.Text())
	}
	return nil
}
