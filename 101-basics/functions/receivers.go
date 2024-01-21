package main

import (
	"fmt"
	"time"
)

type Like struct {
	Date time.Time
}

type User struct {
	Name  string
	Likes []Like
}

func (u *User) Like() {

	like := Like{
		Date: time.Now(),
	}

	u.Likes = append(u.Likes, like)
}

func (u User) LikeCount() int {
	return len(u.Likes)
}

func main() {

	defer fmt.Println("Program sonlandiriliyor.")

	user_1 := User{
		Name: "user_1",
	}

	user_1.Like()
	user_1.Like()
	user_1.Like()

	fmt.Println(user_1.LikeCount())

}
