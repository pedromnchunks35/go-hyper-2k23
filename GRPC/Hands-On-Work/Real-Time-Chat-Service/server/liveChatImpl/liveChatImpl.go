package liveChatImpl

import (
	c "chat/protofiles"
	"context"
	"fmt"
	"io"
	"sync"
)

type ChatServer struct {
	c.UnimplementedChatServer
	Users            []*c.UserData
	UsersMutex       sync.RWMutex
	Connections      map[string]*c.Chat_JoinServer
	ConnectionsMutex sync.RWMutex
}

// ? Function to register a user
func (csv *ChatServer) Register(ctx context.Context, newUser *c.UserData) (*c.Confirmation, error) {
	csv.UsersMutex.Lock()
	defer csv.UsersMutex.Unlock()
	for _, user := range csv.Users {
		if user.Username == newUser.Username {
			return nil, fmt.Errorf("the user is already registered")
		}
	}
	csv.Users = append(csv.Users, newUser)
	conf := &c.Confirmation{}
	conf.Msg = fmt.Sprintf("The user %v got registered", newUser.Username)
	return conf, nil
}

// ? Function to join the channel and send messages
func (csv *ChatServer) Join(stream c.Chat_JoinServer) error {
	//? Get the first message
	firstMessage, err := stream.Recv()
	if err != nil {
		return fmt.Errorf("you need to provide credentials in order to join the channel: %v", err)
	}
	//? LOCK
	csv.UsersMutex.RLock()
	csv.ConnectionsMutex.Lock()
	//? Credentials checking
	userFound := false
	for _, user := range csv.Users {
		if user.Username == firstMessage.Credentials.Username &&
			user.Password == firstMessage.Credentials.Password {
			userFound = true
			//? Add the stream to the array
			csv.Connections[firstMessage.Credentials.Username] = &stream
			break
		}
	}
	//? UNLOCK
	csv.UsersMutex.RUnlock()
	csv.ConnectionsMutex.Unlock()
	//? Case the user has not found return error
	if !userFound {
		return fmt.Errorf("please provide valid credentials")
	}
	message := &c.Message{}
	message.Msg = fmt.Sprintf("The user %v just joined the channel", firstMessage.Credentials.Username)
	message.Sender = "server"
	//? LOCK
	csv.UsersMutex.RLock()
	csv.ConnectionsMutex.RLock()
	//? Tell everyone a given user joined the channel
	for _, user := range csv.Users {
		connection := csv.Connections[user.Username]
		if connection != nil {
			err := (*connection).Send(message)
			if err != nil {
				fmt.Printf("[Server] Some error occurred %v \n", err)
			}
		}
	}
	//? Unlock
	csv.UsersMutex.RUnlock()
	csv.ConnectionsMutex.RUnlock()
	//? Start the listening of streams
	for {
		//? Read next message
		newInfo, err := stream.Recv()
		//? In end of stream end the loop and put the given user without any connection
		if err == io.EOF {
			//? LOCK THE Connections
			csv.ConnectionsMutex.Lock()
			//? Remove the connection
			csv.Connections[firstMessage.Credentials.Username] = nil
			//? Unlock
			csv.ConnectionsMutex.Unlock()
			break
		}
		if err != nil {
			//? LOCK THE Connections
			csv.ConnectionsMutex.Lock()
			//? Remove the connection
			csv.Connections[firstMessage.Credentials.Username] = nil
			//? Unlock
			csv.ConnectionsMutex.Unlock()
			break
		}
		//? Lock
		csv.UsersMutex.RLock()
		csv.ConnectionsMutex.RLock()
		//? Send the message to all the clients or just for one depending of the receiver
		if newInfo.Receiver != "" {
			for _, user := range csv.Users {
				connection := csv.Connections[user.Username]
				if user.Username != firstMessage.Credentials.Username && connection != nil && user.Username == newInfo.Receiver {
					(*connection).Send(newInfo)
					break
				}
			}
		} else {
			for _, user := range csv.Users {
				connection := csv.Connections[user.Username]
				if user.Username != firstMessage.Credentials.Username && connection != nil {
					(*connection).Send(newInfo)
				}
			}
		}
		//? Unlock
		csv.UsersMutex.RUnlock()
		csv.ConnectionsMutex.RUnlock()
	}
	message = &c.Message{}
	message.Msg = fmt.Sprintf("The user %v just left the channel", firstMessage.Credentials.Username)
	message.Sender = "server"
	//? lock
	csv.UsersMutex.RLock()
	csv.ConnectionsMutex.RLock()
	//? Tell everyone a given user left the channel
	for _, user := range csv.Users {
		connection := csv.Connections[user.Username]
		if connection != nil {
			err := (*connection).Send(message)
			if err != nil {
				fmt.Printf("[Server] Some error occurred (someone closed to fast) %v \n", err)
			}
		}
	}
	//? Unlock
	csv.UsersMutex.RUnlock()
	csv.ConnectionsMutex.RUnlock()
	return nil
}
