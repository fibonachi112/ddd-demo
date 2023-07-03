package entity

import (
	"crypto/md5"
	"fmt"
	"time"
)

// Сущность пользователя
type User struct {
	Id        int
	Age       int
	Name      string
	Surname   string
	Email     string
	Password  string
	CreatedAt string
	UpdatedAt string
}

func (u *User) SetPassword(password string) {
	u.Password = fmt.Sprintf("%x", md5.New().Sum([]byte(password)))
}

func (u *User) SetCreatedAtNow() {
	u.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
}
func (u *User) SetUpdatedAtNow() {
	u.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
}

func MakeUser(age int, name, surname, email, password string) User {
	user := User{
		Id:        0,
		Age:       age,
		Name:      name,
		Surname:   surname,
		Email:     email,
		Password:  "",
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt: "",
	}

	user.SetPassword(password)
	return user
}
