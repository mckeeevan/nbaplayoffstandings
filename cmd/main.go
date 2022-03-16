package main

import (
	"fmt"

)

func main() {
	stuff := input("hello")
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



func input(filename string) []Matchup {
    var games []Matchup
	games = append(games, Matchup{team1:"Den", team2: "Mil", team1prob: .9302, team2prob: .1202})
	return games
}