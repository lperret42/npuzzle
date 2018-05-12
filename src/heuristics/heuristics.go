package heuristics

import (
	"lib"
	"math"
)

type Heuristic func(puzzle []int, parent []int,
	goal []int, heuristic float64, last_nb_move int, size int) float64

func DistManhattan(x0 int, y0 int, x1 int, y1 int) float64 {
	return (math.Abs(float64(x1-x0)) + math.Abs(float64(y1-y0)))
}

func DistEuclidean(x0 int, y0 int, x1 int, y1 int) float64 {
	return math.Sqrt(float64((x1-x0)*(x1-x0) + (y1-y0)*(y1-y0)))
}

func indexTwoGoodNums(raw []int, value1 int, value2 int) (int, int) {
	index1 := -1
	index2 := -1
	if value1 != 0 {
		index1 = lib.SliceIntIndex(raw, value1)
	}
	if value2 != 0 {
		index2 = lib.SliceIntIndex(raw, value2)
	}
	return index1, index2
}

func linearConflictRaw(puzzleRaw []int, goalRaw []int, size int) int {
	linearConflictRaw := 0
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if j != i {
				puzzleIndex_1, puzzleIndex_2 := indexTwoGoodNums(puzzleRaw,
					puzzleRaw[i], puzzleRaw[j])
				goalIndex_1, goalIndex_2 := indexTwoGoodNums(goalRaw,
					puzzleRaw[i], puzzleRaw[j])
				if puzzleIndex_1 != -1 && puzzleIndex_2 != -1 &&
					goalIndex_1 != -1 && goalIndex_2 != -1 &&
					(puzzleIndex_1 < puzzleIndex_2 && goalIndex_1 > goalIndex_2 ||
						puzzleIndex_1 > puzzleIndex_2 && goalIndex_1 < goalIndex_2) {
					linearConflictRaw++
				}
			}
		}
	}
	return linearConflictRaw
}

func Manhattan(puzzle []int, parent []int,
	goal []int, heuristic float64, last_nb_move int, size int) float64 {

	var x0, y0, x1, y1, x2, y2 int
	var dist float64
	if last_nb_move == 0 {
		for i := 0; i < len(puzzle); i++ {
			if puzzle[i] != 0 {
				x0, y0 = lib.GetRectCoord(puzzle, size, size, puzzle[i])
				x1, y1 = lib.GetRectCoord(goal, size, size, puzzle[i])
				dist += DistManhattan(x0, y0, x1, y1)
			}
		}
	} else {
		x0, y0 = lib.GetRectCoord(puzzle, size, size, last_nb_move)
		x1, y1 = lib.GetRectCoord(goal, size, size, last_nb_move)
		x2, y2 = lib.GetRectCoord(parent, size, size, last_nb_move)
		if DistManhattan(x1, y1, x2, y2) > DistManhattan(x0, y0, x1, y1) {
			dist = heuristic - 1
		} else {
			dist = heuristic + 1
		}
	}
	return (dist)
}

func Euclidean(puzzle []int, parent []int,
	goal []int, heuristic float64, last_nb_move int, size int) float64 {
	var x0, y0, x1, y1 int
	var dist float64
	for i := 0; i < len(puzzle); i++ {
		if puzzle[i] != 0 {
			x0, y0 = lib.GetRectCoord(puzzle, size, size, puzzle[i])
			x1, y1 = lib.GetRectCoord(goal, size, size, puzzle[i])
			dist += DistEuclidean(x0, y0, x1, y1)
		}
	}
	return (dist)
}

func Dijkstra(puzzle []int, parent []int,
	goal []int, heuristic float64, last_nb_move int, size int) float64 {
	return 0
}

func LinearConflict(puzzle []int, parent []int,
	goal []int, heuristic float64, last_nb_move int, size int) float64 {

	puzzleRaw := []int{}
	goalRaw := []int{}
	linearConflict := 0
	i := 0

	for size+i <= len(puzzle) {
		linearConflict += linearConflictRaw(puzzle[i:size+i], goal[i:size+i], size)
		i += size
	}
	for i := 0; i < size; i++ {
		for j := 0; j < len(puzzle); {
			puzzleRaw = append(puzzleRaw, puzzle[i+j])
			goalRaw = append(goalRaw, goal[i+j])
			j += size
		}
		linearConflict += linearConflictRaw(puzzleRaw, goalRaw, size)
		puzzleRaw = puzzleRaw[:0]
		goalRaw = goalRaw[:0]
	}
	return (float64(linearConflict))
}

func ManhattanLinearConflict(puzzle []int, parent []int,
	goal []int, heuristic float64, last_nb_move int, size int) float64 {

	return Manhattan(puzzle, parent, goal, heuristic, last_nb_move, size) +
		LinearConflict(puzzle, parent, goal, heuristic, last_nb_move, size)
}

func MisplacedTiles(puzzle []int, parent []int,
	goal []int, heuristic float64, last_nb_move int, size int) float64 {

	misplaced := 0
	for i := 0; i < len(puzzle); i++ {
		puzzleIndex:= lib.SliceIntIndex(puzzle, puzzle[i])
		goalIndex := lib.SliceIntIndex(goal, puzzle[i])
		if puzzleIndex != goalIndex && puzzle[i] != 0 {
			misplaced++
		}
	}
	return float64(misplaced)
}
