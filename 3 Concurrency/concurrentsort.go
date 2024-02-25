package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func main() {
	// Prompt the user to input a series of integers
	nums := takeInput()

	// Partition the array into 4 parts
	numPartitions := 4
	partitionSize := (len(nums) + numPartitions - 1) / numPartitions
	partitions := make([][]int, numPartitions)
	for i := 0; i < numPartitions; i++ {
		start := i * partitionSize
		end := (i + 1) * partitionSize
		if end > len(nums) {
			end = len(nums)
		}
		partitions[i] = nums[start:end]
	}

	// Spawn a goroutine to sort each partition
	var wg sync.WaitGroup
	wg.Add(numPartitions)
	for i := 0; i < numPartitions; i++ {
		go func(i int) {
			defer wg.Done()
			sort.Ints(partitions[i])
			fmt.Printf("Partition %d: %v\n", i+1, partitions[i])
		}(i)
	}
	wg.Wait()

	// Merge the sorted subarrays into one large sorted array
	var sortedNums []int
	for _, partition := range partitions {
		sortedNums = merge(sortedNums, partition)
	}

	// Print the entire sorted list
	fmt.Printf("Sorted list: %v\n", sortedNums)
}

// func takes [] and merges them
func merge(left, right []int) []int {
	merged := make([]int, 0, len(left)+len(right))
	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			return append(merged, right...)
		}
		if len(right) == 0 {
			return append(merged, left...)
		}
		if left[0] <= right[0] {
			merged = append(merged, left[0])
			left = left[1:]
		} else {
			merged = append(merged, right[0])
			right = right[1:]
		}
	}
	return merged
}

// func takes input and creates a []int
func takeInput() []int {
	var int_slc []int
take_input:
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("enter a space-separated sequence of integers.")
	scanner.Scan()
	inp := strings.ToLower(scanner.Text())
	req := strings.Split(inp, " ")
	for _, v := range req {
		i, err := strconv.Atoi(v)
		if err != nil {
			log.Printf("[Error] Invalid input!")
			goto take_input
		}
		int_slc = append(int_slc, i)
	}
	fmt.Println("The unsorted slice of integers is: ", int_slc)
	return int_slc
}
