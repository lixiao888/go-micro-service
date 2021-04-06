package handler

import (
	"context"
	"user/domain/model"
	"user/domain/service"

	user "user/proto/user"
)

type User struct{
	UserDataService service.IUserDataService
}

//注册
func (u *User)Register(ctx context.Context,userRegisterRequest *user.UserRegisterRequest,userRegisterResponse *user.UserRegisterResponse) error {
	userRegister := &model.User{
		UserName:    userRegisterRequest.UserName ,
		FirstName:   userRegisterRequest.FirstName,
		HashPassword:userRegisterRequest.Pwd,
	}
	_,err :=u.UserDataService.AddUser(userRegister)
	if err !=nil {
		return err
	}
	userRegisterResponse.Message = "添加成功"
	return nil
}
