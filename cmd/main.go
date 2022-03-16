package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	data := grabCSV()
	fmt.Println(data)
	convert(data)
}

type Matchup struct {
	team1      string
	team2      string
	team1prob  float64
	team2prob  float64
	team1score float64
	team2score float64
}

type Matchuptemp struct {
	team1      string
	team2      string
	team1prob  string
	team2prob  string
	team1score string
	team2score string
}

type teamStanding struct {
	teamName    string
	teamLoses   float64
	teamWins    float64
	winningPerc float64
}

func convert(data []Matchuptemp) ([]Matchup, []teamStanding) {
	var matchups []Matchup
	var standings []teamStanding

	for _, v := range data {
		if v.team1score == "" {
			// the game hasn't been played

		}
		if v.team1score != "" {
			// the game has been played

		}
	}

	return matchups, standings

}

func grabCSV() []Matchuptemp {

	records, err := readData("nba_elo_latest.csv")

	if err != nil {
		log.Fatal(err)
	}
	var matchups []Matchuptemp

	for _, record := range records {

		temp := Matchuptemp{
			team1:      record[4],
			team2:      record[5],
			team1prob:  record[20],
			team2prob:  record[21],
			team1score: record[22],
			team2score: record[23],
		}

		matchups = append(matchups, temp)

	}

	return matchups

}

func readData(fileName string) ([][]string, error) {

	f, err := os.Open(fileName)

	if err != nil {
		return [][]string{}, err
	}

	defer f.Close()

	r := csv.NewReader(f)

	// skip first line
	if _, err := r.Read(); err != nil {
		return [][]string{}, err
	}

	records, err := r.ReadAll()

	if err != nil {
		return [][]string{}, err
	}

	return records, nil
}
