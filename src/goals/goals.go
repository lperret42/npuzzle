package goals

func GetSnake(size int) []int {
	goal := make([]int, size*size)
	depl := "right"
	var right_max, down_max, left_max, up_max int = 0, 0, 0, 1
	i := 0
	j := 0
	nb := 1
	for nb < len(goal) {
		goal[i*size+j] = nb
		if depl == "right" && !(j < size-1-right_max) {
			depl = "down"
			right_max++
		} else if depl == "down" && !(i < size-1-down_max) {
			depl = "left"
			down_max++
		} else if depl == "left" && !(j > left_max) {
			depl = "up"
			left_max++
		} else if depl == "up" && !(i > up_max) {
			depl = "right"
			up_max++
		}
		if depl == "right" {
			j++
		} else if depl == "down" {
			i++
		} else if depl == "left" {
			j--
		} else if depl == "up" {
			i--
		}
		nb++
	}
	return goal
}

func GetLinear(size int) []int {
	goal := make([]int, size*size)
	for i := 0; i < len(goal)-1; i++ {
		goal[i] = i + 1
	}
	goal[len(goal)-1] = 0
	return goal
}
