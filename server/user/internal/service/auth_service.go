package service

type IAuthService interface {
}

type AuthService struct{}

func NewAuthService() IAuthService {
	return &AuthService{}
}
