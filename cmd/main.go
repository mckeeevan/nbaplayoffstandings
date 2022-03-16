package main

import (
	"fmt"

	"github.com/mckeeevan/nbaplayoffstandings/pkg/inputCSV"
)

func main() {
	stuff := inputCSV.input("hello")
	fmt.Println(stuff)
}


type team struct{
	name string
	wins int8
	loses int8
}

type Matchup struct{
	team1 string
	team2 string
	team1prob float64
	team2prob float64
}
