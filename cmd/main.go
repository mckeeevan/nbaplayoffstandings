package main

import "fmt"

func main() {
	fmt.Println("Hello, 世界")
}


type team struct{
	name
	wins
	loses
}

type matchup struct{
	team1
	team2
	team1prob
	team2prob
}
