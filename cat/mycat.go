package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Начало работы программы")
	checkArgs()
	fmt.Println("Конец работы программы")
}

func checkArgs() {
	numberOfArgs := len(os.Args)
	if numberOfArgs == 2 {
		readText()
	} else {
		fmt.Println("Укажите только один файл")
	}
}

// :3
func readText() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Проверьте правильность указанных файлов")
	} else {
		fmt.Println("Чтение файла...")
		Data := bufio.NewScanner(file)
		for Data.Scan() {
			fmt.Printf("%s\n", Data.Text())
		}
	}
}
