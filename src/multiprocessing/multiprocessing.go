package multiprocessing

import (
	"displaying"
	"fmt"
	"npuzzle"
	"runtime"
	"sync"
)

func RunMultipleAlgo(npuzzle npuzzle.Npuzzle) {
	npuzzle2 := npuzzle
	npuzzle3 := npuzzle
	npuzzle4 := npuzzle
	runtime.GOMAXPROCS(4)
	var wg sync.WaitGroup
	wg.Add(1)
	finish := false
	go func() {
		npuzzle.Solve("A", "manhattan", 3)
		if !finish {
			finish = true
			displaying.DisplaySol(npuzzle, true)
			fmt.Println("\nThe winner is A with manhattan")
			wg.Done()
		}
	}()
	go func() {
		npuzzle2.Solve("A", "manhattan_lc", 2)
		if !finish {
			finish = true
			displaying.DisplaySol(npuzzle2, true)
			fmt.Println("\nThe winner is A with manhattan_lc")
			wg.Done()
		}
	}()
	go func() {
		npuzzle3.Solve("A", "euclidean", 4)

		if !finish {
			finish = true
			displaying.DisplaySol(npuzzle3, true)
			fmt.Println("yo")
			fmt.Println("\nThe winner is A with euclidean")
			wg.Done()
		}
	}()
	go func() {
		npuzzle4.Solve("A", "linear_conflict", 5)
		if !finish {
			finish = true
			displaying.DisplaySol(npuzzle4, true)
			fmt.Println("\nThe winner is A with linear_conflict")
			wg.Done()
		}
	}()
	wg.Wait()
}
