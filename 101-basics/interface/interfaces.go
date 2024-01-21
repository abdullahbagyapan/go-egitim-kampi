package main

import (
	"time"
)

type User interface {
	Like()
	LikeCount() int
}

type user struct {
	Name  string
	Likes []Like
}

type Like struct {
	Date time.Time
}

func (u *user) Like() {

	like := Like{
		Date: time.Now(),
	}

	u.Likes = append(u.Likes, like)
}

func (u user) LikeCount() int {
	return len(u.Likes)
}

func main() {

	//...
}
