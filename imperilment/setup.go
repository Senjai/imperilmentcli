package imperilment

type Engine struct {
    BaseUrl string
    OverallLeaderBoardUrl string
}

func New(baseurl string) *Engine {
    return &Engine{
        BaseUrl: baseurl,
        OverallLeaderBoardUrl: baseurl + "/leader_board/overall.json",
    }
}

