package main

import (
    "fmt"
    "github.com/Senjai/imperilmentcli/imperilment"
)

func main() {
    engine := imperilment.New("http://imperilment.freerunningtech.com")

    board, ok := engine.GetOverallLeaderBoard()
    if !ok {
        fmt.Println("Could not get leader board from ", engine.OverallLeaderBoardUrl, ".")
        return
    }

    fmt.Println("Top 5 results by trophies...")
    for _, entry := range board[:5] {
        fmt.Printf(entry.String())
    }

    board, ok = engine.GetMoneyLeaderBoard()
    if !ok {
        fmt.Println("Could not get leader board from ", engine.MoneyLeaderBoardUrl, ".")
    }

    fmt.Println("Top 5 results by money...")
    for _, entry := range board[:5] {
        fmt.Printf(entry.String())
    }
}
