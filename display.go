package main

import (
    "fmt"
    "github.com/Senjai/imperilmentcli/imperilment"
    "github.com/daviddengcn/go-colortext"
    "math"
    "strconv"
    "strings"
)

var (
    engine =  imperilment.New("http://imperilment.freerunningtech.com")
)

type BoardLengths struct {
    Board int
    Total int
    Name int
}

func DisplayOverallLeaderBoard() {
    defer ct.ResetColor()

    board, ok := engine.GetOverallLeaderBoard()
    if !ok { ct.ChangeColor(ct.Red, false, ct.None, false)
        fmt.Println("Could not get leader board from: ", engine.OverallLeaderBoardUrl)
        return
    }

    lengths := GetBoardLengths(board)

    fmt.Printf("%+v\n", lengths)

    fmt.Println("%s", strings.Repeat("-", lengths.Board)
    fmt.Printf("-%s")
    for _, entry := range board {
    }
}

func GetBoardLengths(board []imperilment.LeaderBoardEntry) *BoardLengths {
    // unchanging
    goldLen := 7
    silverLen := 9
    bronzeLen := 9

    var nameLen float64
    var totalLen float64
    for _, entry := range board {
        nameLen = math.Max(nameLen, float64(len(entry.User.FullName())))
        totalLen = math.Max(totalLen, float64(len(strconv.Itoa(entry.Total))))
    }

    boardLen := 8 + int(nameLen) + int(totalLen) + goldLen + silverLen + bronzeLen

    return &BoardLengths {
        boardLen,
        int(totalLen),
        int(nameLen),
    }
}


