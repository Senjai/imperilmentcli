package imperilment

import (
    "encoding/json"
    "net/http"
    "io/ioutil"
    "log"
)

func (e *Engine) GetOverallLeaderBoard() (board []LeaderBoardEntry, ok bool) {
    return e.getBoard(e.OverallLeaderBoardUrl)
}

func (e *Engine) GetMoneyLeaderBoard() (board []LeaderBoardEntry, ok bool) {
    return e.getBoard(e.MoneyLeaderBoardUrl)
}

func (e *Engine) getBoard(url string) (board []LeaderBoardEntry, ok bool) {
    response, err := http.Get(url)
    if err != nil {
        log.Fatal(err)
        return board, false
    }
    defer response.Body.Close()

    body, _ := ioutil.ReadAll(response.Body)

    json.Unmarshal(body, &board)
    return board, true
}

