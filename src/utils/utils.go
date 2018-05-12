package utils

import (
	"fmt"
	"lib"
)

func GetNbInversions(grid []int, size int) int {
	nb_inversions := 0
	for pos := 0; pos < len(grid)-1; pos++ {
		if grid[pos] == 0 {
			continue
		}
		for i := pos + 1; i < len(grid); i++ {
			if grid[i] == 0 {
				continue
			}
			if grid[i] < grid[pos] {
				nb_inversions++
			}
		}
	}
	if size%2 == 0 {
		nb_inversions += lib.SliceIntIndex(grid, 0) / size
	}
	return (nb_inversions)
}

func PrintGrid(grid []int, size int) {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			nb := grid[i*size+j]
			if size <= 3 {
				if nb == 0 {
					fmt.Printf("%s ", " ")
				} else {
					fmt.Printf("%d ", nb)
				}
			} else if size <= 10 {
				if nb == 0 {
					fmt.Printf("%2s ", " ")
				} else {
					fmt.Printf("%2d ", nb)
				}
			} else {
				if nb == 0 {
					fmt.Printf("%3s ", " ")
				} else {
					fmt.Printf("%3d ", nb)
				}
			}
		}
		fmt.Printf("\n")
	}
}

func GetSwappedSlice(slice []int, index_0 int, index_1 int) []int {
	swapped_slice := make([]int, len(slice))
	copy(swapped_slice, slice)
	swapped_slice[index_0] = slice[index_1]
	swapped_slice[index_1] = slice[index_0]
	return (swapped_slice)
}
