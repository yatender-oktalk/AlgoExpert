package main

import "fmt"

var competitions = [][]string{
	{"HTML", "C#"},
	{"C#", "Python"},
	{"Python", "NExt"},
	{"NExt", "HTML"},
}

var results = []int{1, 0, 1, 0}

const HOME_TEAM_WIN = 1

func main() {
	fmt.Println(GetTournamentWinner(competitions, results))

}

func GetTournamentWinner(competitions [][]string, results []int) string {
	resultsVal := make([]int, len(results))
	for i := 0; i < len(results); i++ {
		if results[i] != 0 {
			resultsVal[i] += 3
		} else {
			remainder := (i + 1) % len(results)
			resultsVal[remainder] += 3
		}
	}

	return GetMax(competitions, resultsVal)
}

func GetTournamentWinnerNew(competitions [][]string, results []int) string {
	CurrentBestTeam := ""
	scores := map[string]int{CurrentBestTeam: 0}

	for idx, competition := range competitions {
		result := results[idx]
		homeTeam, awayTeam := competition[0], competition[1]

		winningTeam := awayTeam

		if result == HOME_TEAM_WIN {
			winningTeam = homeTeam
		}

		updateScores(winningTeam, 3, scores)

		if scores[winningTeam] > scores[CurrentBestTeam] {
			CurrentBestTeam = winningTeam
		}
	}
	return CurrentBestTeam
}

// This will return index
func GetMax(competitions [][]string, arr []int) string {
	maxIndex := 0

	for i := 0; i < len(arr); i++ {
		if arr[i] > arr[maxIndex] {
			maxIndex = i
		}
	}

	return competitions[maxIndex][0]
}

func updateScores(team string, points int, scores map[string]int) {
	scores[team] += points
}
