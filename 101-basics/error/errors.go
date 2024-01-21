package main

import (
	"errors"
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

func (u *User) Like() error {

	if u.LikeCount() == 2 {
		return errors.New("Error !! Max like count is equal 2")
	}

	like := Like{
		Date: time.Now(),
	}

	u.Likes = append(u.Likes, like)
	return nil
}

func (u User) LikeCount() int {
	return len(u.Likes)
}

func main() {

	defer fmt.Println("Program sonlandiriliyor.")

	user_1 := User{
		Name: "user_1",
	}

	for {
		err := user_1.Like()

		if err != nil {
			fmt.Println(err.Error())
			break
		}
	}
}
