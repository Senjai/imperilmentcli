package imperilment

import (
    "encoding/json"
    "net/http"
    "io/ioutil"
    "log"
)

func (e *Engine) GetLeaderBoard() (board []LeaderBoardEntry, ok bool) {
    response, err := http.Get(e.OverallLeaderBoardUrl)
    if err != nil {
        log.Fatal(err)
        return board, false
    }
    defer response.Body.Close()

    body, _ := ioutil.ReadAll(response.Body)

    json.Unmarshal(body, &board)
    return board, true
}
