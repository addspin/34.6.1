package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Использование: program input.txt output.txt")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	// Читаем содержимое входного файла
	content, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Ошибка при чтении файла: %v\n", err)
		return
	}

	// Создаем регулярное выражение для поиска математических выражений
	re := regexp.MustCompile(`(\d+)([\+\-])(\d+)=\?`)

	// Открываем файл для записи результатов
	file, err := os.Create(outputFile)
	if err != nil {
		fmt.Printf("Ошибка при создании файла: %v\n", err)
		return
	}
	defer file.Close()

	// Создаем буферизированный writer
	writer := bufio.NewWriter(file)

	// Обрабатываем каждую строку входного файла
	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		matches := re.FindStringSubmatch(line)
		if len(matches) > 0 {
			num1, _ := strconv.Atoi(matches[1])
			operator := matches[2]
			num2, _ := strconv.Atoi(matches[3])

			var result int
			switch operator {
			case "+":
				result = num1 + num2
			case "-":
				result = num1 - num2
			}

			// Записываем результат в буфер
			fmt.Fprintf(writer, "%d%s%d=%d\n", num1, operator, num2, result)
		}
	}

	// Сбрасываем буфер в файл
	writer.Flush()
}
