package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"sync"
	"time"
	u "user-auth/protofiles"

	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/examples/data"
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

var (
	tls      = flag.Bool("tls", false, "Connection use tls if true, else plain TCP")
	certFile = flag.String("cert_file", "", "TLS cert file relative path")
	keyFile  = flag.String("key_file", "", "The TLS key file")
)

func main() {
	flag.Parse()

	bind, err := net.Listen("tcp", "localhost:2000")
	if err != nil {
		log.Fatalf("Something went wrong with the binding %v", err)
	}
	//? Opts
	var opts []grpc.ServerOption
	//? Check if tls is requested as flag
	if *tls {
		//? Check the given attributes
		if *certFile == "" {
			*certFile = data.Path("x509/server_cert.pem")
		}
		if *keyFile == "" {
			*keyFile = data.Path("x509/server_key.pem")
		}
		//? Create credentials from the certificates
		creds, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
		if err != nil {
			log.Fatalf("Failed to generate credentials: %v", err)
		}
		//? add the tls credentials
		opts = []grpc.ServerOption{grpc.Creds(creds)}
	}
	grpc := grpc.NewServer(opts...)
	u.RegisterUserAuthServer(grpc, &UserAuthService{Users: []*u.Credentials{}, JwtSecret: []byte{1, 3, 6, 100, 200, 30, 2, 8}})
	log.Printf("The server will now listen into %v", bind.Addr())
	err = grpc.Serve(bind)
	if err != nil {
		log.Fatalf("Something went wrong serving the server %v", err)
	}
}
