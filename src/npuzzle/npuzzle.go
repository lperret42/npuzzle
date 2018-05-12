package npuzzle

import (
	"fmt"
	"goals"
	"heuristics"
	"lib"
	"math"
	"node"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"utils"
)

const (
	COMMENT_CHAR = "#"
)

type Npuzzle struct {
	Raw_text    string
	Size        int
	Goal        []int
	Puzzle      []int
	MaxOpenSet  int
	MaxInMemory int
	Sol         [][]int
}

func (npuzzle Npuzzle) PrintInitialMap() {
	fmt.Println(npuzzle.Raw_text)
}

func (npuzzle Npuzzle) checkUnicity() int {
	encountered := map[int]bool{}
	for i := range npuzzle.Puzzle {
		if encountered[npuzzle.Puzzle[i]] == true {
			return (-1)
		} else {
			encountered[npuzzle.Puzzle[i]] = true
		}
	}
	return (0)
}

func (npuzzle *Npuzzle) parsePuzzleLine(line string) int {
	nb_str := strings.Fields(line)
	if len(nb_str) != npuzzle.Size {
		return (-1)
	}
	for i := 0; i < npuzzle.Size; i++ {
		nb, err := strconv.Atoi(nb_str[i])
		if err != nil || !(nb >= 0 && nb < npuzzle.Size*npuzzle.Size) {
			return (-1)
		}
		npuzzle.Puzzle = append(npuzzle.Puzzle, nb)
	}
	return (0)
}

func (npuzzle *Npuzzle) Parse() int {
	lines := strings.Split(npuzzle.Raw_text, "\n")
	lib.RemoveComment(lines, COMMENT_CHAR)
	file_len := len(lines)
	num_puzzle_line := 0
	for i := 0; i < file_len; i++ {
		if len(lines[i]) == 0 {
			continue
		} else if npuzzle.Size == 0 {
			size, err := strconv.Atoi(lines[i])
			if err != nil || size < 3 {
				return (-1)
			}
			npuzzle.Size = size
		} else {
			if num_puzzle_line >= npuzzle.Size {
				return (-1)
			}
			if npuzzle.parsePuzzleLine(lines[i]) != 0 {
				return (-1)
			}
			num_puzzle_line += 1
		}
	}
	if npuzzle.checkUnicity() != 0 {
		return (-1)
	}
	return (0)
}

func (npuzzle *Npuzzle) GetGoal(goal_shape string) bool {
	if goal_shape == "snake" {
		npuzzle.Goal = goals.GetSnake(npuzzle.Size)
	} else if goal_shape == "linear" {
		npuzzle.Goal = goals.GetLinear(npuzzle.Size)
	} else {
		return false
	}
	return true
}

func (npuzzle Npuzzle) CheckSolvable() bool {
	puzzle_nb_inversions := utils.GetNbInversions(npuzzle.Puzzle, npuzzle.Size)
	goal_nb_inversions := utils.GetNbInversions(npuzzle.Goal, npuzzle.Size)
	if puzzle_nb_inversions%2 != goal_nb_inversions%2 {
		return false
	}
	return true
}

func Reverse(puzzles [][]int) {
	nb_puzzles := len(puzzles)
	for i := 0; i < nb_puzzles/2; i++ {
		tmp := puzzles[nb_puzzles-1-i]
		puzzles[nb_puzzles-1-i] = puzzles[i]
		puzzles[i] = tmp
	}
}

func (npuzzle *Npuzzle) getSol(end node.Node) {
	for end.Parent != nil {
		npuzzle.Sol = append(npuzzle.Sol, end.Puzzle)
		end = *end.Parent
	}
	npuzzle.Sol = append(npuzzle.Sol, end.Puzzle)
	lib.ReverseIntSliceArray(npuzzle.Sol)
}

func (npuzzle *Npuzzle) PrintInfos() {
	fmt.Println("Total number of states ever selected in the \"opened\" set:",
		npuzzle.MaxOpenSet)
	fmt.Println("Maximum number of states ever represented in memory",
		"at the same time during the search:", npuzzle.MaxInMemory)
	fmt.Println("Number of moves required to transition from the initial state",
		"to the final state, according to the search:", len(npuzzle.Sol)-1)
}

func (npuzzle *Npuzzle) PrintSol() {
	if len(npuzzle.Sol) == 0 {
		return
	}
	fmt.Println("Solution:\n")
	for i := 0; i < len(npuzzle.Sol); i++ {
		utils.PrintGrid(npuzzle.Sol[i], npuzzle.Size)
		fmt.Println("")
	}
	/*
		fmt.Println("Total number of states ever selected in the \"opened\" set:",
			npuzzle.MaxOpenSet)
		fmt.Println("Maximum number of states ever represented in memory",
			"at the same time during the search:", npuzzle.MaxInMemory)
		fmt.Println("Number of moves required to transition from the initial state",
			"to the final state, according to the search:", len(npuzzle.Sol)-1)
	*/
}

func (npuzzle *Npuzzle) iDASearch(root node.Node, h heuristics.Heuristic,
	threshold float64) float64 {
	var cost float64
	var nbCut int
	openList := []node.Node{}
	nextThreshold := math.MaxFloat64

	openList = append(openList, root)
	for len(openList) != 0 {
		npuzzle.MaxOpenSet = lib.Max(npuzzle.MaxOpenSet, len(openList))
		npuzzle.MaxInMemory = lib.Max(npuzzle.MaxInMemory, len(openList)+nbCut)
		process := openList[0]
		if lib.DeepEqual(process.Puzzle, npuzzle.Goal) {
			npuzzle.getSol(process)
			return -1
		}
		openList = openList[1:]
		nbCut++
		cost = process.GetCost()
		if cost > threshold {
			nbCut--
			nextThreshold = math.Min(cost, nextThreshold)
		} else {
			openList = append(openList, process.GetSons(h, npuzzle.Goal)...)
		}
	}
	return nextThreshold
}

func (npuzzle *Npuzzle) iDAStar(h heuristics.Heuristic, boost float64) {
	threshold := h(npuzzle.Puzzle, nil, npuzzle.Goal, 0, 0, npuzzle.Size)
	root := node.Node{npuzzle.Size, npuzzle.Puzzle, 0, threshold, boost, nil}
	MULTITHREADING := false
	if MULTITHREADING {
		nb_cores := runtime.GOMAXPROCS(8)
		var wg sync.WaitGroup
		finish := false
		wg.Add(1)
		for i := 0; i < nb_cores; i++ {
			go func() {
				for {
					if npuzzle.iDASearch(root, h, threshold) <= 0 {
						break
					}
					threshold += 2
				}
				if !finish {
					finish = true
					wg.Done()
				}
			}()
			threshold += 2
		}
		wg.Wait()
	} else {
		for {
			threshold = npuzzle.iDASearch(root, h, threshold)
			if threshold <= 0 {
				return
			}
		}
	}
}

func (npuzzle *Npuzzle) aStar(h heuristics.Heuristic, boost float64) {
	openList := map[float64][]node.Node{}  // cost -> nodes
	closedList := make(map[string]float64) //string_puzzle -> depth
	var first_heuristic, best_cost float64
	var process_string, son_string string
	var sons []node.Node
	var nbElemOpenList int
	var nbElemClosedList int

	first_heuristic = h(npuzzle.Puzzle, nil, npuzzle.Goal, 0, 0, npuzzle.Size)
	root := node.Node{npuzzle.Size, npuzzle.Puzzle, 0, first_heuristic, boost, nil}
	best_cost = root.GetCost()
	openList[best_cost] = append(openList[best_cost], root)
	for {
		nbElemOpenList = GetNbElemInMap(openList)
		nbElemClosedList = len(closedList)
		npuzzle.MaxOpenSet = lib.Max(npuzzle.MaxOpenSet, nbElemOpenList)
		npuzzle.MaxInMemory = lib.Max(npuzzle.MaxInMemory, nbElemOpenList+nbElemClosedList)
		best_cost = MinKey(openList)
		process := openList[best_cost][0]
		if lib.DeepEqual(process.Puzzle, npuzzle.Goal) {
			npuzzle.getSol(process)
			break
		}
		openList[best_cost] = openList[best_cost][1:]
		if len(openList[best_cost]) == 0 {
			delete(openList, best_cost)
		}
		sons = process.GetSons(h, npuzzle.Goal)
		for i := 0; i < len(sons); i++ {
			son_string = lib.SliceIntToString(sons[i].Puzzle)
			if depth, exist := closedList[son_string]; exist {
				if sons[i].Depth >= depth {
					continue
				}
			}
			son_cost := sons[i].GetCost()
			isInOpenCostInf := false
			for cost, nodes := range openList {
				if cost >= son_cost {
					continue
				}
				if node.IsInNodes(nodes, sons[i].Puzzle) {
					isInOpenCostInf = true
					break
				}
			}
			if isInOpenCostInf {
				continue
			}
			openList[son_cost] = append(openList[son_cost], sons[i])
		}
		process_string = lib.SliceIntToString(process.Puzzle)
		closedList[process_string] = process.Depth
	}
}

func (npuzzle *Npuzzle) Solve(algo string, heuristic string, boost float64) bool {
	if algo == "A" {
		if heuristic == "manhattan" {
			npuzzle.aStar(heuristics.Manhattan, boost)
		} else if heuristic == "euclidean" {
			npuzzle.aStar(heuristics.Euclidean, boost)
		} else if heuristic == "dijkstra" {
			npuzzle.aStar(heuristics.Dijkstra, boost)
		} else if heuristic == "linear_conflict" {
			npuzzle.aStar(heuristics.LinearConflict, boost)
		} else if heuristic == "manhattan_lc" {
			npuzzle.aStar(heuristics.ManhattanLinearConflict, boost)
		} else {
			return false
		}
	} else if algo == "IDA" {
		if heuristic == "manhattan" {
			npuzzle.iDAStar(heuristics.Manhattan, boost)
		} else if heuristic == "euclidean" {
			npuzzle.iDAStar(heuristics.Euclidean, boost)
		} else if heuristic == "dijkstra" {
			npuzzle.iDAStar(heuristics.Dijkstra, boost)
		} else if heuristic == "linear_conflict" {
			npuzzle.iDAStar(heuristics.LinearConflict, boost)
		} else if heuristic == "manhattan_lc" {
			npuzzle.iDAStar(heuristics.ManhattanLinearConflict, boost)
		} else {
			return false
		}
	}
	return true
}

func MinKey(m map[float64][]node.Node) float64 {
	var k, min float64
	min = math.MaxFloat64
	for k, _ = range m {
		if k < min {
			min = k
		}
	}
	return min
}

func GetNbElemInMap(m map[float64][]node.Node) int {
	sum := 0
	for _, value := range m {
		sum += len(value)
	}
	return (sum)
}
