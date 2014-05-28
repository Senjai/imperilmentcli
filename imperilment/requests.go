package imperilment

import (
    "encoding/json"
    "net/http"
    "io/ioutil"
    "log"
)

func (e *Engine) GetOverallLeaderBoard() (board []LeaderBoardEntry, ok bool) {
    ok = e.getBoard(e.OverallLeaderBoardUrl, &board)
    return board, ok
}

func (e *Engine) GetMoneyLeaderBoard() (board []LeaderBoardEntry, ok bool) {
    ok = e.getBoard(e.MoneyLeaderBoardUrl, &board)
    return board, ok
}

func (e *Engine) getBoard(url string, board interface{}) (ok bool) {
    response, err := http.Get(url)
    if err != nil {
        log.Fatal(err)
        return false
    }
    defer response.Body.Close()

    body, _ := ioutil.ReadAll(response.Body)

    json.Unmarshal(body, board)
    return true
}

