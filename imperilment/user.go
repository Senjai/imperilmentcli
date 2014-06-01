package imperilment

import "fmt"

type User struct {
    Email string
    First_Name string
    Last_Name string
}

func (u *User) FullName() string {
    return fmt.Sprintf("%s, %s", u.First_Name, u.Last_Name)
}

