package main

import "fmt"
import "io/ioutil"
import "bytes"
import "strings"

type User struct {
	ID   string
	Name string
}

var ALL_USERS []User

func ReadUser() {
	users, _ := ioutil.ReadFile("user")
	buffer := bytes.NewBuffer(users)
	for {
		line, e := buffer.ReadString('\n')
		if e != nil {
			break
		}
		line = strings.Trim(line, "\n")
		user := strings.Split(line, "|")
		fmt.Println(len(line), line)
		if len(line) <= 0 {
			continue
		}
		ALL_USERS = append(ALL_USERS, User{ID: user[0], Name: user[1]})
	}
	for _, user := range ALL_USERS {
		fmt.Println(user)
	}
}
