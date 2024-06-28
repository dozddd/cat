package main

import (
	"bytes"
	"errors"
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
		writer         io.Writer
		expectedOutput string
		expectedError  error
	}{
		{
			name:           "Чтение файла",
			reader:         strings.NewReader("Первая строка\nВторая строка\nТретья строка"),
			writer:         &bytes.Buffer{},
			expectedOutput: "Первая строка\nВторая строка\nТретья строка",
			expectedError:  nil,
		},
		{
			name:           "Одна строка",
			reader:         strings.NewReader("Первая строка"),
			writer:         &bytes.Buffer{},
			expectedOutput: "Первая строка",
			expectedError:  nil,
		},
		{
			name:           "Пустой файл",
			reader:         strings.NewReader(""),
			writer:         &bytes.Buffer{},
			expectedOutput: "",
			expectedError:  nil,
		},
		{
			name:           "nil reader",
			reader:         nil,
			writer:         &bytes.Buffer{},
			expectedOutput: "",
			expectedError:  ReaderIsNilErr,
		},
		{
			name:           "nil writer",
			reader:         strings.NewReader("Первая строка"),
			writer:         nil,
			expectedOutput: "",
			expectedError:  WriterIsNilErr,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			err := readText(testCase.reader, testCase.writer)
			if testCase.expectedError != nil {
				if !errors.Is(err, testCase.expectedError) {
					t.Fail()
				}
			}

			if testCase.writer != nil {
				actualOutput := testCase.writer.(*bytes.Buffer).String()
				if actualOutput != testCase.expectedOutput {
					t.Fail()
				}
			}
		})
	}
}
