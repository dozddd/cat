package main

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

func TestCheckArgs_Success(t *testing.T) {
	tCase := []struct {
		name           string
		inputArgs      []string
		expectedResult bool
	}{
		{
			name:           "Не передано путь до файла",
			inputArgs:      []string{"./mycat"},
			expectedResult: false,
		},
		{
			name:           "Передано более одного пути до файла",
			inputArgs:      []string{"./mycat", "path/to/file1.txt", "path/to/file2.txt"},
			expectedResult: false,
		},
		{
			name:           "Валидные аргументы",
			inputArgs:      []string{"./mycat", "path/to/file1.txt"},
			expectedResult: true,
		},
	}

	for _, tCase := range tCase {
		t.Run(tCase.name, func(t *testing.T) {
			if actual := checkArgs(tCase.inputArgs); actual != tCase.expectedResult {
				t.Fail()
			}
		})
	}
}

func TestReadText(t *testing.T) {
	testCases := []struct {
		name           string
		reader         io.Reader
		expectedOutput string
		expectedError  error
	}{
		{
			name:           "Чтение файла",
			reader:         strings.NewReader("Первая строка\nВторая строка\nТретья строка"),
			expectedOutput: "Первая строка\nВторая строка\nТретья строка",
			expectedError:  nil,
		},
		{
			name:           "Одна строка",
			reader:         strings.NewReader("Первая строка"),
			expectedOutput: "Первая строка",
			expectedError:  nil,
		},
		{
			name:           "Пустой файл",
			reader:         strings.NewReader(""),
			expectedOutput: "",
			expectedError:  nil,
		},
	}

	for _, testCases := range testCases {
		t.Run(testCases.name, func(t *testing.T) {
			var buf bytes.Buffer
			actual := readText(testCases.reader, &buf)

			if actual != testCases.expectedError {
				t.Fail()
			}

			if buf.String() != testCases.expectedOutput {
				t.Fail()
			}
		})
	}
}
