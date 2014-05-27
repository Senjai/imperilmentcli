package main

import (
    "fmt"
    "encoding/json"
    "net/http"
    "io/ioutil"
    "log"
)

type LeaderBoard struct {
    User User
    Total int
    First int
    Second int
    Third int
}

func (b *LeaderBoard) String() string {
    return fmt.Sprintf(
        "%s, %s: %v Gold Trophies, %v Silver Trophies, %v Bronze Trophies, Total: %v\n",
        b.User.First_Name,
        b.User.Last_Name,
        b.First,
        b.Second,
        b.Third,
        b.Total,
    )
}

type User struct {
    Email string
    First_Name string
    Last_Name string
}

func main() {
    response, err := http.Get("http://imperilment.freerunningtech.com/leader_board/overall.json")
    if err != nil {
        log.Fatal(err)
        return
    }
    defer response.Body.Close()
    body, _ := ioutil.ReadAll(response.Body)

    var board []LeaderBoard
    json.Unmarshal(body, &board)

    for _, entry := range board {
        fmt.Printf(entry.String())
    }
}
