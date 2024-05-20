package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Чтение входных данных
func readInput() ([][]int, int, error) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return nil, 0, err
	}
	containers := make([][]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		line := scanner.Text()
		values := strings.Fields(line)
		containers[i] = make([]int, n)
		for j := 0; j < n; j++ {
			containers[i][j], err = strconv.Atoi(values[j])
			if err != nil {
				return nil, 0, err
			}
		}
	}
	return containers, n, nil
}

func canSortContainers(containers [][]int, n int) bool {
	// Суммирование количества шаров каждого цвета
	colorSum := make([]int, n)
	containerSum := make([]int, n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			colorSum[j] += containers[i][j]
			containerSum[i] += containers[i][j]
		}
	}

	// Сортировка для сравнения
	sort.Ints(colorSum)
	sort.Ints(containerSum)

	// Проверка равенства сумм
	for i := 0; i < n; i++ {
		if colorSum[i] != containerSum[i] {
			return false
		}
	}
	return true
}

func main() {
	containers, n, err := readInput()
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	if canSortContainers(containers, n) {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}
