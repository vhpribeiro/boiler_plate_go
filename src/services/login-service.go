package services

import (
	"time"

	"boiler_plate.com/src/repositorys"
	"boiler_plate.com/src/services/dtos"
	"boiler_plate.com/src/utils/errors"
	"github.com/dgrijalva/jwt-go"
)

type ILoginService interface {
	Login(loginDto dtos.LoginDto) (map[string]string, error)
}

type loginService struct {
	userRepository repositorys.IUserRepository
}

func (l *loginService) Login(loginDto dtos.LoginDto) (map[string]string, error) {

	user := l.userRepository.GetUser(loginDto.Username, loginDto.Password)
	if user.ID == 0 {
		return nil, &errors.InternalError{Message: "User not found!"}
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = "Jon Doe"
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	// The signing string should be secret (a generated UUID works too)
	tokenGenerated, err := token.SignedString([]byte("secret"))
	if err != nil {
		return nil, err
	}
	mapResult := map[string]string{
		"token": tokenGenerated,
	}
	return mapResult, nil
}

func NewLoginService(userRepository repositorys.IUserRepository) ILoginService {
	return &loginService{
		userRepository: userRepository,
	}
}
