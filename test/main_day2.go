package main

import (
	"fmt"
	"ws/output"
	"ws/unionfind"
)

func main() {
	ds := unionfind.NewDisjointSet(6)
	ds.Union(0, 1)
	ds.Union(1, 2)
	ds.Union(3, 4)

	output.PrintDisjointSet(ds)

	ds.Union(2, 3)
	output.PrintDisjointSet(ds)
	fmt.Println("Parent:", ds.GetParent())
	fmt.Println("Rank:  ", ds.GetRank())
}
