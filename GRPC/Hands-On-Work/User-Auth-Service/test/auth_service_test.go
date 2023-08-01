package test

import (
	"context"
	"fmt"
	"strings"
	"testing"
	u "user-auth/protofiles"
)

var token string

func Test_Register(t *testing.T) {
	req := &u.Credentials{}
	req.Username = "teste"
	req.Password = "12341234"
	confirmation, err := Client.Register(context.Background(), req)
	if err != nil {
		t.Fatalf("cannot throw error, register is valid %v", err)
	}
	if !strings.Contains(confirmation.Msg, fmt.Sprintf("The user %v got successfully registed", req.Username)) {
		t.Fatalf("should container success message")
	}
	_, err = Client.Register(context.Background(), req)
	if !strings.Contains(err.Error(), "the user already exists") {
		t.Fatalf("user should already exist")
	}
}

func Test_Login_Invalid(t *testing.T) {
	req := &u.Credentials{}
	req.Username = "teste12345"
	req.Password = "12341234"
	_, err := Client.Login(context.Background(), req)
	if !strings.Contains(err.Error(), "user not found") {
		t.Fatalf("should throw a error saying user not found")
	}
	req.Username = "teste"
	req.Password = "1234"
	_, err = Client.Login(context.Background(), req)
	if !strings.Contains(err.Error(), "user not found") {
		t.Fatalf("should throw a error saying user not found")
	}
}

func Test_Login(t *testing.T) {
	req := &u.Credentials{}
	req.Username = "teste"
	req.Password = "12341234"
	confirmation, err := Client.Login(context.Background(), req)
	if err != nil {
		t.Fatalf("should not throw error")
	}
	if !strings.Contains(confirmation.Msg, "Login was sucessfull") {
		t.Fatalf("should return success message")
	}
	if len(confirmation.Jwt) == 0 {
		t.Fatalf("should retrieve a token")
	}
	token = confirmation.Jwt
}

func Test_HelloWorld_Invalid(t *testing.T) {
	req := &u.Jwt{}
	req.Msg = "sishdgbsiug"
	_, err := Client.HelloWorld(context.Background(), req)
	if !strings.Contains(err.Error(), "token is not valid") {
		t.Fatalf("should return that the token is invalid")
	}
}

func Test_HelloWorld(t *testing.T) {
	req := &u.Jwt{}
	req.Msg = token
	confirmation, err := Client.HelloWorld(context.Background(), req)
	if err != nil {
		t.Fatalf("should not return a error")
	}
	if !strings.Contains(confirmation.Msg, "Hello World") {
		t.Fatalf("should return hello world")
	}
}
