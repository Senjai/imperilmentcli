package main

import (
    "fmt"
    "github.com/Senjai/imperilmentcli/imperilment"
)

func main() {
    engine := imperilment.New("http://imperilment.freerunningtech.com")

    board, ok := engine.GetLeaderBoard()
    if !ok {
        fmt.Println("Could not get leader board from ", engine.OverallLeaderBoardUrl, ".")
        return
    }

    for _, entry := range board {
        fmt.Printf(entry.String())
    }
}
