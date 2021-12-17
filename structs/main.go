package main

import (
	"fmt"
	"log"
	"os"
)

type User struct {
	name  string
	email string
}

func (u User) Name() string {
	return u.name
}

func (u User) SetName(name string) {
	u.name = name
}

func (u *User) Email() string {
	return u.email
}

func (u *User) SetEmail(email string) {
	u.email = email
}

type Job struct {
	Name string
	*log.Logger
}

func (j *Job) Println(v ...interface{}) {
	j.Logger.Println("prefix-", v)
}

func main() {
	user := User{}
	user.SetName("Alex")
	user.SetEmail("ryazanov@go.com")

	fmt.Println(user.Name())
	fmt.Println(user.Email())
	fmt.Println("-------------")

	job := &Job{"qwe", log.New(os.Stderr, "Job: ", log.Ldate)}
	job.Println("qwe")
}
