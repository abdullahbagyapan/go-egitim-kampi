package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type User struct {
	Id           int       `json:"-"`
	Name         string    `json:"Name"`
	Followers    []User    `json:"Followers,omitempty"`
	RegisterDate time.Time `json:"-"`
}

func main() {

	user_1 := User{
		Id:           1,
		Name:         "user_1",
		RegisterDate: time.Now(),
		Followers: []User{
			{
				Id:           2,
				Name:         "user_2",
				RegisterDate: time.Now(),
			},
			{
				Id:           3,
				Name:         "user_3",
				RegisterDate: time.Now(),
			},
		},
	}

	arr, _ := json.Marshal(user_1)
	fmt.Println(string(arr))
}
