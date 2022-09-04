package main

import (
	"fmt"

	d "github.com/alaleks/delislice"
)

type MyStructure struct {
	Val int
}

func main() {
	sl := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sl, err := d.DelItemInd(8, sl)
	fmt.Println(sl, err)
	fmt.Println("-----------------------------")

	sl2 := []int{1, 2, 2, 3, 4, 5, 6, 7, 7}
	sl2, err = d.DelItemInd(2, sl2)
	fmt.Println(sl2, err)
	fmt.Println("-----------------------------")

	sl3 := []string{"1", "2", "2", "3", "4", "5", "6"}
	sl3, err = d.DelItemVal("2", sl3)
	fmt.Println(sl3, err)
	fmt.Println("-----------------------------")

	sl4 := []MyStructure{
		{Val: 1},
		{Val: 1},
		{Val: 2},
		{Val: 3}}
	sl4, err = d.DelItemVal(MyStructure{Val: 1}, sl4)
	fmt.Println(sl4, err)
}
