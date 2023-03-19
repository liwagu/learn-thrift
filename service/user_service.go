// 编写service/user_service.go以实现Thrift接口。

package service

import (
	"context"
	"learn-thrift/pkg/thrift/example"
)

type UserServiceImpl struct{}

func (u *UserServiceImpl) GetUserProfile(ctx context.Context, userId int32) (*example.UserProfile, error) {
	// 在这里实现具体的业务逻辑
	userProfile := &example.UserProfile{
		ID:    userId,
		Name:  "John Doe",
		Email: "john.doe@example.com",
	}
	return userProfile, nil
}
