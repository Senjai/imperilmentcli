package imperilment

type Engine struct {
    BaseUrl string
    OverallLeaderBoardUrl string
    MoneyLeaderBoardUrl string
    GameLeaderBoardUrl string
}

func New(baseurl string) *Engine {
    return &Engine{
        BaseUrl: baseurl,
        OverallLeaderBoardUrl: baseurl + "/leader_board/overall.json",
        MoneyLeaderBoardUrl: baseurl + "/leader_board/money.json",
        GameLeaderBoardUrl: baseurl + "/leader_board.json",
    }
}

