package imperilment

import "fmt"

type GameLeaderBoardEntry struct {
    User User
    Results []string
    Total int
    Position int
}

func (b *GameLeaderBoardEntry) ResultRatio() (correct int, numanswers int) {
    var results []string
    for _, result := range b.Results {
        if result != "unavailable" {
            results = append(results, result)
            if result == "correct" {
                correct++
            }
        }
    }
    return correct, len(results)
}


func (b *GameLeaderBoardEntry) String() string {
    correct, answers := b.ResultRatio()
    return fmt.Sprintf(
        "%s, %s: Position: %v, Ratio: %v/%v, Total: %v",
        b.User.First_Name,
        b.User.Last_Name,
        correct,
        answers,
        b.Total,
    )
}


