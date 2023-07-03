package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	DiFabric "repo/internal/DI/fabric"
	"repo/internal/http/requests"
)



func UserCreateOne(ctx *fiber.Ctx) error {
	di := DiFabric.Make()
	user, err := requests.MakeUserEntityFromRequest(ctx)
	if err != nil {
		fmt.Println(fmt.Sprintf("requsest err : %s", err))
	}

	err = di.UserRepository.SaveOne(&user)
	if err != nil {
		fmt.Println(fmt.Sprintf("save err : %s", err))
	}

	return ctx.JSON(map[string]interface{}{
		"user_id": user.Id,
	})
}


func UserGetOne(ctx *fiber.Ctx) error {
	di := DiFabric.Make()
	user_id, err := ctx.ParamsInt("user_id",-1)

	if err != nil || user_id <0{
		return ctx.JSON(map[string]interface{}{
			"code": "error",
		})
	}

	user, err := di.UserRepository.GetOne(user_id)

	return ctx.JSON(map[string]interface{}{
		"user_id": user.Id,
		"name": fmt.Sprintf( "%s %s", user.Name, user.Surname) ,
		"email": user.Email,
	})
}
