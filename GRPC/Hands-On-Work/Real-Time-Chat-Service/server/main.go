package main

import (
	c "chat/protofiles"
	"context"
	"fmt"
	"sync"
)

type ChatServer struct {
	c.UnimplementedChatServer
	users            []*c.UserData
	usersMutex       sync.RWMutex
	connections      map[string]*c.Chat_JoinServer
	connectionsMutex sync.RWMutex
}

// ? Function to register a user
func (csv *ChatServer) Register(ctx context.Context, newUser *c.UserData) (*c.Confirmation, error) {
	csv.usersMutex.Lock()
	defer csv.usersMutex.Unlock()
	for _, user := range csv.users {
		if user.Username == newUser.Username {
			return nil, fmt.Errorf("the user is already registered")
		}
	}
	csv.users = append(csv.users, newUser)
	conf := &c.Confirmation{}
	conf.Msg = fmt.Sprintf("The user %v got registered", newUser.Username)
	return conf, nil
}
