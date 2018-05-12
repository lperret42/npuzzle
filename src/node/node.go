package node

import (
	"fmt"
	"heuristics"
	"lib"
	"math"
	"utils"
)

type Node struct {
	Size      int
	Puzzle    []int
	Depth     float64
	Heuristic float64
	Boost     float64
	Parent    *Node
}

func (node *Node) puzzleLikeGrandPa() bool {
	if node.Parent == nil || node.Parent.Parent == nil {
		return false
	} else if !lib.DeepEqual(node.Puzzle, node.Parent.Parent.Puzzle) {
		return false
	}
	return true
}

func (node *Node) getNode(swap_1 int, swap_2 int, h heuristics.Heuristic,
	boost float64, goal []int) *Node {
	new_node := Node{Size: node.Size, Depth: node.Depth + 1, Parent: node}
	new_node.Puzzle = utils.GetSwappedSlice(node.Puzzle, swap_1, swap_2)
	if new_node.puzzleLikeGrandPa() {
		return nil
	}
	new_node.Heuristic = h(new_node.Puzzle, node.Puzzle, goal,
		node.Heuristic, node.Puzzle[swap_2], node.Size)
	new_node.Boost = boost
	return &new_node
}

func (parent *Node) GetSons(h heuristics.Heuristic, goal []int) []Node {
	var nodes []Node
	var tmp *Node
	x_0, y_0 := lib.GetRectCoord(parent.Puzzle, parent.Size, parent.Size, 0)
	index_0 := x_0*parent.Size + y_0
	if x_0 > 0 {
		if tmp = parent.getNode(index_0, index_0-parent.Size, h,
			parent.Boost, goal); tmp != nil {
			nodes = append(nodes, *tmp)
		}
	}
	if x_0 < parent.Size-1 {
		if tmp = parent.getNode(index_0, index_0+parent.Size, h,
			parent.Boost, goal); tmp != nil {
			nodes = append(nodes, *tmp)
		}
	}
	if y_0 > 0 {
		if tmp = parent.getNode(index_0, index_0-1, h,
			parent.Boost, goal); tmp != nil {
			nodes = append(nodes, *tmp)
		}
	}
	if y_0 < parent.Size-1 {
		if tmp = parent.getNode(index_0, index_0+1, h,
			parent.Boost, goal); tmp != nil {
			nodes = append(nodes, *tmp)
		}
	}
	return nodes
}

func (node *Node) GetCost() float64 {
	return node.Depth + math.Pow(node.Heuristic, node.Boost)
}

func (node *Node) Print() {
	fmt.Println("size: ", node.Size)
	fmt.Println("Depth: ", node.Depth)
	fmt.Println("heuristic: ", node.Heuristic)
	fmt.Println("parent: ", node.Parent)
	fmt.Println("Puzzle: ")
	utils.PrintGrid(node.Puzzle, node.Size)
}

func IsInNodes(nodes []Node, puzzle []int) bool {
	for i := 0; i < len(nodes); i++ {
		if lib.DeepEqual(nodes[i].Puzzle, puzzle) {
			return true
		}
	}
	return false
}
