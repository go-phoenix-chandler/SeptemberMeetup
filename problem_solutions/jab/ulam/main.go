package main

import (
	"flag"
	"fmt"
	"github.com/kloster/ulam/prime"
	"strconv"
)

var (
	start int
	level int
	graph bool
)

func init() {
	flag.IntVar(&start, "start", 1, "Which number to start with")
	flag.IntVar(&level, "level", 12, "How many Ulam levels to go")
	flag.BoolVar(&graph, "graph", false, "Whether to graph the prime numbers as '#' rather than numbers. Default to false")
	flag.Parse()
}

func main() {

	Ulamit(start, level)
}

func Ulamit(start, level int) {
	counter := start
	point := []int{level, level}

	grid := [][]int{}
	for x := 0; x < level+level+1; x++ {
		grid = append(grid, fillSlice(level))
	}

	grid[point[0]][point[1]] = counter //middle
	counter++

	for x := 1; x < level+1; x++ {

		// right 1
		point = []int{point[0], point[1] + 1}
		grid[point[0]][point[1]] = counter
		counter++

		// go up (y axis) x * 2 - 1
		moveUp := x*2 - 1
		for up := 1; up <= moveUp; up++ {
			point = []int{point[0] - 1, point[1]}
			grid[point[0]][point[1]] = counter
			counter++
		}

		moveLeft := x + x
		for left := 1; left <= moveLeft; left++ {
			point = []int{point[0], point[1] - 1}
			grid[point[0]][point[1]] = counter
			counter++
		}

		moveDown := x + x
		for down := 1; down <= moveDown; down++ {
			point = []int{point[0] + 1, point[1]}
			grid[point[0]][point[1]] = counter
			counter++
		}

		moveRight := x + x
		for right := 1; right <= moveRight; right++ {
			point = []int{point[0], point[1] + 1}
			grid[point[0]][point[1]] = counter
			counter++
		}

	}

	printGrid(grid)
}

func fillSlice(amount int) (sl []int) {
	for x := 0; x < amount+amount+1; x++ {
		sl = append(sl, 0)
	}

	return
}

func printGrid(grid [][]int) {
	for _, x := range grid {
		for _, i := range x {
			if prime.Is(i) {
				if !graph {
					fmt.Print(" # ")
				} else {
					fmt.Print(pad4(i) + " ")
				}

			} else {
				if !graph {
					if i == 1 {
						fmt.Print(" + ")
					} else {
						fmt.Print(" . ")
					}

				} else {
					fmt.Print(".... ")
				}
			}
		}
		fmt.Println("")
	}
}

func pad4(num int) string {
	strNum := strconv.Itoa(num)

	if len(strNum) == 1 {
		strNum = ".." + strNum + "."
	} else if len(strNum) == 2 {
		strNum = "." + strNum + "."
	} else if len(strNum) == 3 {
		strNum = "." + strNum
	}

	return strNum
}
