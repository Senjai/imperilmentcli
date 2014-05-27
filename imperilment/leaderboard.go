package imperilment

import "fmt"

type LeaderBoardEntry struct {
    User User
    Total int
    First int
    Second int
    Third int
}

type User struct {
    Email string
    First_Name string
    Last_Name string
}

func (b *LeaderBoardEntry) String() string {
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
