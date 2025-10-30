package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}
	// 先对区间按照起始位置进行排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	merged := [][]int{}
	currentInterval := intervals[0]

	for _, interval := range intervals[1:] {
		if interval[0] <= currentInterval[1] {
			// 有重叠，合并区间
			currentInterval[1] = max(currentInterval[1], interval[1])
		} else {
			// 无重叠，添加当前区间到结果中
			merged = append(merged, currentInterval)
			currentInterval = interval
		}
	}
	// 添加最后一个区间
	merged = append(merged, currentInterval)
	return merged
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	result := merge(intervals)
	fmt.Println("Merged Intervals:", result)
}
