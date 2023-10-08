package main

import (
	"bytes"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

func handleInput(file string) [][]int {
	fileContent, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("Error reading file", err)
		return nil
	}

	var data [][]int
	var group []int

	lines := bytes.Split(fileContent, []byte("\n"))

	for _, line := range lines {
		if len(line) == 0 {
			if len(group) > 0 {
				data = append(data, group)
				group = make([]int, 0)
			}
			continue
		}

		num, err := strconv.Atoi(string(line))
		if err != nil {
			fmt.Println("Conversion error: ",err)
			continue
		}

		group = append(group, num)

	}

	// Append the last group if it's not empty
	if len(group) > 0 {
		data = append(data, group)
	}

	return data
}



func sumCalories(calories []int) int {
	sum := 0
	for _, calorie := range calories {
		sum += calorie
	}
	return sum
}

type MinHeap []int

func (h MinHeap) Len() int              { return len(h) }
func (h MinHeap) Less(i, j int) bool    { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface {}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop()  interface {} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0:n-1]
	return x
}

func main() {
	file := "day1-input.txt"
	data := handleInput(file)
	
		var elfs MinHeap
		heap.Init(&elfs)

		for _, group := range data {
			total := sumCalories(group)
			
			if elfs.Len() < 3 {
				heap.Push(&elfs,total)
			} else {
				min := elfs[0]
				if total > min {
					heap.Pop(&elfs)
					heap.Push(&elfs, total)
				}
			}
		}

		topThree := make([]int, 3)

		for i := 2; i >= 0; i-- {
			topThree[i] = heap.Pop(&elfs).(int)
		}

		total := sumCalories(topThree)
		
		fmt.Printf("%v Top three total calories: %d\n", topThree, total)
}