package inject

import (
	userCreateUseCase "github.com/Poul-george/go-api/api/core/usecase/api/user/create"
	userDetailUseCase "github.com/Poul-george/go-api/api/core/usecase/api/user/detail"
	userListUseCase "github.com/Poul-george/go-api/api/core/usecase/api/user/list"
	userCreate "github.com/Poul-george/go-api/api/presentation/appapi/controller/user/create"
	userDetail "github.com/Poul-george/go-api/api/presentation/appapi/controller/user/detail"
	userList "github.com/Poul-george/go-api/api/presentation/appapi/controller/user/list"
	"github.com/labstack/echo/v4"
)

func (i *Injector) V1UserListController() echo.HandlerFunc {
	return newHandlerFunc(userList.Controller{
		UseCase: *userListUseCase.NewUseCase(i.userRepository()),
	}.Get)
}

func (i *Injector) V1UserDetailController() echo.HandlerFunc {
	return newHandlerFunc(userDetail.Controller{
		UseCase: *userDetailUseCase.NewUseCase(i.userRepository()),
	}.Get)
}

func (i *Injector) V1UserCreateController() echo.HandlerFunc {
	return newHandlerFunc(userCreate.Controller{
		UseCase: *userCreateUseCase.NewUseCase(i.userRepository()),
	}.Post)
}
