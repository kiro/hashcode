package main

import (
	"google/hashcode/task"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		panic("Expected one argument file name.")
	}
	task.Solve(os.Args[1])
}
