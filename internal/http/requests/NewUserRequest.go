package requests

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"repo/internal/domain/entity"
	"time"
)

type NewUserRequest struct {
	Name     string `json:"name,omitempty"`
	Surname  string `json:"surname,omitempty"`
	Age      int    `json:"age,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

func MakeUserEntityFromRequest(ctx *fiber.Ctx) (entity.User, error){
	newUserRequest := new(NewUserRequest)

	if err := ctx.BodyParser(newUserRequest); err != nil {
		return entity.User{}, errors.New(fmt.Sprintf("error : %s", err))
	}

	user := entity.User{
		Age:      newUserRequest.Age,
		Name:     newUserRequest.Name,
		Surname:  newUserRequest.Surname,
		Email:    newUserRequest.Email,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}
	user.SetPassword(newUserRequest.Password)
	return user, nil

}
