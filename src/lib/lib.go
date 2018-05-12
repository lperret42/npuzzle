package lib

import (
	"bytes"
	"strconv"
	"strings"
)

func RemoveComment(lines []string, comment_char string) {
	for i := 0; i < len(lines); i++ {
		lines[i] = strings.Split(lines[i], comment_char)[0]
	}
}

func SliceIntIndex(slice []int, value int) int { // return first occurrence
	for i := 0; i < len(slice); i++ { //of value in slice
		if slice[i] == value {
			return (i)
		}
	}
	return (-1)
}

func GetRectCoord(rect []int, x_size int, y_size int, value int) (int, int) {
	index := SliceIntIndex(rect, value)
	return index / x_size, index % y_size
}

func ReverseIntSliceArray(array [][]int) {
	nb := len(array)
	for i := 0; i < nb/2; i++ {
		tmp := array[nb-1-i]
		array[nb-1-i] = array[i]
		array[i] = tmp
	}
}

func Abs(x int) int {
	if x >= 0 {
		return x
	} else {
		return -x
	}
}

func Max(x, y int) int {
	if x >= y {
		return x
	} else {
		return y
	}
}

func Min(x, y int) int {
	if x <= y {
		return x
	} else {
		return y
	}
}

func DeepEqual(a []int, b []int) bool {
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func SliceIntToString(slice []int) string {
	var buffer bytes.Buffer

	for i := 0; i < len(slice); i++ {
		buffer.WriteString(strconv.Itoa(slice[i]))
	}
	return buffer.String()
}
