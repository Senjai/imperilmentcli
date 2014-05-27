package main

import (
    "fmt"
    "github.com/Senjai/imperilmentcli/imperilment"
)

func main() {
    url := "http://imperilment.freerunningtech.com/leader_board/overall.json"
    board, ok := imperilment.GetLeaderBoard(url)
    if !ok {
        fmt.Println("Could not get leader board from ", url, ".")
        return
    }
    for _, entry := range board {
        fmt.Printf(entry.String())
    }
}
