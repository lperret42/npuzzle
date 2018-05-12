package main

import (
	"displaying"
	"flag"
	"fmt"
	"io/ioutil"
	"multiprocessing"
	"npuzzle"
	"os"
)

func parse_args(show_map *bool, display *string, boost *float64,
	goal, algo, heuristic *string, multi_processing *bool) {
	flag.BoolVar(show_map, "show_map", false, "print initial map")
	flag.StringVar(display, "display", "verbose", "verbose or user")
	flag.StringVar(goal, "goal", "snake", "choose shape of goal: snake or linear")
	flag.StringVar(algo, "algo", "A", "choose an algo: A or IDA")
	flag.StringVar(heuristic, "h", "manhattan", "choose an heuristic: "+
		"manhattan, euclidean, linear_conflict, manhattan_lc or dijkstra")
	flag.Float64Var(boost, "boost", 1, "choose a boost level: >= 1")
	flag.BoolVar(multi_processing, "multi_processing", false, "run multiple algo at the same time."+
		"If true, flags algo and h are ignored")
	flag.Parse()
}

func main() {
	var show_map bool
	var display string
	var boost float64
	var goal, algo, heuristic string
	var multi_processing bool
	parse_args(&show_map, &display, &boost, &goal, &algo, &heuristic, &multi_processing)
	if boost < 1 {
		fmt.Println("Error: boost must be a real number greater than 1")
		return
	}
	files := flag.Args()
	if len(files) < 1 {
		os.Stderr.WriteString("Error: not exactly one input file\n")
		return
	}
	file := files[0]
	raw_text, err := ioutil.ReadFile(file)
	if err != nil {
		os.Stderr.WriteString("Error: can't open input file\n")
		return
	}
	npuzzle := npuzzle.Npuzzle{Raw_text: string(raw_text)}
	if show_map {
		npuzzle.PrintInitialMap()
	}
	if npuzzle.Parse() != 0 {
		fmt.Println("Parsing error")
		return
	}
	if !npuzzle.GetGoal(goal) {
		fmt.Println("error: doesn't handle this kind of goal")
		return
	}
	if !npuzzle.CheckSolvable() {
		fmt.Println("This puzzle is unsolvable")
		return
	}

	if multi_processing {
		multiprocessing.RunMultipleAlgo(npuzzle)
	} else {
		//fmt.Println("\nSolvable:", npuzzle.CheckSolvable(), "\n")
		if !npuzzle.Solve(algo, heuristic, boost) {
			fmt.Println("error: algo", algo, "with heuristic", heuristic, "is not a good mix")
			return
		}
		if display == "user" {
			displaying.DisplaySol(npuzzle, show_map)
		} else if display == "verbose" {
			npuzzle.PrintSol()
			npuzzle.PrintInfos()
		} else {
			npuzzle.PrintInfos()
		}
	}
}
