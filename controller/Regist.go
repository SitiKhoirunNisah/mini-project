package controller

import (
	"model/model"
	"model/service"
)

type AuthController struct {
	authService service.AuthService
	userView    view.UserView
}

func NewAuthController(authService service.AuthService, userView view.UserView) *AuthController {
	return &AuthController{
		authService: authService,
		userView:    userView,
	}
}

func (c *AuthController) RegisterUser(name, email, password string) {
	newUser := model.User{
		Name:     name,
		Email:    email,
		Password: password,
	}
	if err := c.authService.RegisterUser(newUser); err != nil {
		c.userView.ShowRegistrationError(err.Error())
	} else {
		c.userView.ShowRegistrationSuccessMessage()
	}
}

func (c *AuthController) LoginUser(email, password string) {
	user, err := c.authService.AuthenticateUser(email, password)
	if err != nil {
		c.userView.ShowLoginError(err.Error())
	} else {
		c.userView.ShowLoginSuccessMessage(user.Name)
	}
}

func (c *AuthController) LogoutUser() {
	c.authService.LogoutUser()
	c.userView.ShowLogoutMessage()
}
