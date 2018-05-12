package displaying

import (
	"fmt"
	"npuzzle"
	"os"
	"os/exec"
	"utils"
)

func refresh(npuzzle npuzzle.Npuzzle, num_puzzle int, show_map bool) {
	fmt.Printf("\033[2J\033[1;1H")
	if show_map {
		npuzzle.PrintInitialMap()
	}
	fmt.Printf("Initial state: \n")
	utils.PrintGrid(npuzzle.Sol[0], npuzzle.Size)
	fmt.Printf("\nGoal: \n")
	utils.PrintGrid(npuzzle.Goal, npuzzle.Size)
	fmt.Printf("\n")
	npuzzle.PrintInfos()
	fmt.Printf("\n")
	fmt.Printf("Move number %d\n\n", num_puzzle)
	utils.PrintGrid(npuzzle.Sol[num_puzzle], npuzzle.Size)
	fmt.Printf("\nPress \"q\" to quit\n")
}

func DisplaySol(npuzzle npuzzle.Npuzzle, show_map bool) {
	if len(npuzzle.Sol) == 0 {
		return
	}
	exec.Command("stty", "-f", "/dev/tty", "cbreak", "min", "1").Run()
	exec.Command("stty", "-f", "/dev/tty", "-echo").Run()
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
	defer func() {
		exec.Command("stty", "-f", "/dev/tty", "echo").Run()
		exec.Command("stty", "-F", "/dev/tty", "echo").Run()
	}()

	i := 0
	refresh(npuzzle, i, show_map)
	nb_puzzle := len(npuzzle.Sol)
	var b []byte = make([]byte, 1)
	for {
		os.Stdin.Read(b)
		keyboard_nb := b[0]
		if keyboard_nb == 113 {
			break
		} else if keyboard_nb == 68 {
			if i > 0 {
				i--
				refresh(npuzzle, i, show_map)
			}
		} else if keyboard_nb == 67 {
			if i < nb_puzzle-1 {
				i++
				refresh(npuzzle, i, show_map)
			}
		}
	}
}
