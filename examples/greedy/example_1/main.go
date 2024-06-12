package main

import (
	"fmt"
	"sort"
)

// Найти максимальное количество непересекающихся интервалов.
// Входные данные: [{1, 3}, {2, 4}, {3, 5}, {0, 6}, {5, 7}, {8, 9}, {5, 9}]
// Вывод программы: 4

func main() {
	intervals := []Interval{
		{start: 1, end: 3},
		{start: 2, end: 4},
		{start: 3, end: 5},
		{start: 0, end: 6},
		{start: 5, end: 7},
		{start: 8, end: 9},
		{start: 5, end: 9},
	}

	result := maxNonOverlappingIntervals(intervals)
	fmt.Println("Result:", result)
}

type Interval struct {
	start, end int
}

// Функция для сортировки интервалов по времени окончания
func sortByEnd(intervals []Interval) {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].end < intervals[j].end
	})
}

// Функция для нахождения максимального количества непересекающихся интервалов
func maxNonOverlappingIntervals(intervals []Interval) int {
	if len(intervals) == 0 {
		return 0
	}

	// Сортируем интервалы по времени окончания
	sortByEnd(intervals)

	count := 1
	lastEnd := intervals[0].end

	for _, interval := range intervals[1:] {
		if interval.start >= lastEnd {
			count++
			lastEnd = interval.end
		}
	}

	return count
}
