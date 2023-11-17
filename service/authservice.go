package service

import "model/model"

type AuthService struct {
	// Implementasi sesuai kebutuhan
}

func (s *AuthService) RegisterUser(user model.User) error {
	// Implementasi registrasi pengguna
}

func (s *AuthService) AuthenticateUser(email, password string) (*model.User, error) {
	// Implementasi otentikasi pengguna
}

func (s *AuthService) LogoutUser() {
	// Implementasi logout pengguna
}
