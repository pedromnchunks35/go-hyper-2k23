package authServiceImpl

import (
	"context"
	"fmt"
	"sync"
	"time"
	u "user-auth/protofiles"

	"github.com/dgrijalva/jwt-go"
)

type UserAuthService struct {
	*u.UnimplementedUserAuthServer
	Users      []*u.Credentials
	UsersMutex sync.Mutex
	JwtSecret  []byte
	JwtMutex   sync.Mutex
}

func (uas *UserAuthService) Register(ctx context.Context, credentials *u.Credentials) (*u.Confirmation, error) {
	uas.UsersMutex.Lock()
	defer uas.UsersMutex.Unlock()
	for _, user := range uas.Users {
		if user.Username == credentials.Username {
			return nil, fmt.Errorf("the user already exists")
		}
	}
	uas.Users = append(uas.Users, credentials)
	confirmation := &u.Confirmation{}
	confirmation.Msg = fmt.Sprintf("The user %v got successfully registed", credentials.Username)
	return confirmation, nil
}

func (uas *UserAuthService) generateJwt(username string) (string, error) {
	uas.JwtMutex.Lock()
	defer uas.JwtMutex.Unlock()
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), //Make it expire in 24h
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(uas.JwtSecret)
	if err != nil {
		return "", fmt.Errorf("something went wrong in the token generation: %v", err)
	}
	return tokenString, nil
}

func (uas *UserAuthService) verifyJwt(token string) error {
	uas.JwtMutex.Lock()
	defer uas.JwtMutex.Unlock()
	tokenNew, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return uas.JwtSecret, nil
	})
	if err != nil {
		return err
	}
	if _, ok := tokenNew.Claims.(jwt.MapClaims); ok && tokenNew.Valid {
		return nil
	}
	return fmt.Errorf("invalid token")
}

func (uas *UserAuthService) Login(ctx context.Context, credentials *u.Credentials) (*u.ConfirmationPlus, error) {
	uas.UsersMutex.Lock()
	defer uas.UsersMutex.Unlock()
	userFound := false
	for _, user := range uas.Users {
		if user.Username == credentials.Username && user.Password == credentials.Password {
			userFound = true
			break
		}
	}
	if !userFound {
		return nil, fmt.Errorf("user not found")
	}
	jwtToken, err := uas.generateJwt(credentials.Username)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	confirmation := &u.ConfirmationPlus{}
	confirmation.Jwt = jwtToken
	confirmation.Msg = "Login was sucessfull"
	return confirmation, nil
}

func (uas *UserAuthService) HelloWorld(ctx context.Context, jwt *u.Jwt) (*u.Confirmation, error) {
	err := uas.verifyJwt(jwt.Msg)
	if err != nil {
		return nil, fmt.Errorf("token is not valid")
	}
	confirmation := &u.Confirmation{}
	confirmation.Msg = "Hello World"
	return confirmation, nil
}
