package user

import (
	crud "firego/app/user/crud"
)

type UserController struct {
	user_crud crud.UserCRUD
}

func NewUserController() UserController {
	return UserController{user_crud: crud.NewUserCRUD()}
}
