package main

import (
	c "chat/protofiles"
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"sync"

	"google.golang.org/grpc"
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

// ? Function to join the channel and send messages
func (csv *ChatServer) Join(stream c.Chat_JoinServer) error {
	//? Get the first message
	firstMessage, err := stream.Recv()
	if err != nil {
		return fmt.Errorf("you need to provide credentials in order to join the channel: %v", err)
	}
	//? LOCK
	csv.usersMutex.RLock()
	csv.connectionsMutex.Lock()
	//? Credentials checking
	userFound := false
	for _, user := range csv.users {
		if user.Username == firstMessage.Credentials.Username &&
			user.Password == firstMessage.Credentials.Password {
			userFound = true
			//? Add the stream to the array
			csv.connections[firstMessage.Credentials.Username] = &stream
			break
		}
	}
	//? UNLOCK
	csv.usersMutex.RUnlock()
	csv.connectionsMutex.Unlock()
	//? Case the user has not found return error
	if !userFound {
		return fmt.Errorf("please provide valid credentials")
	}
	message := &c.Message{}
	message.Msg = fmt.Sprintf("The user %v just joined the channel", firstMessage.Credentials.Username)
	message.Sender = "server"
	//? LOCK
	csv.usersMutex.RLock()
	csv.connectionsMutex.RLock()
	//? Tell everyone a given user joined the channel
	for _, user := range csv.users {
		connection := csv.connections[user.Username]
		if connection != nil {
			err := (*connection).Send(message)
			if err != nil {
				fmt.Printf("[Server] Some error occurred %v \n", err)
			}
		}
	}
	//? Unlock
	csv.usersMutex.RUnlock()
	csv.connectionsMutex.RUnlock()
	//? Start the listening of streams
	for {
		//? Read next message
		newInfo, err := stream.Recv()
		//? In end of stream end the loop and put the given user without any connection
		if err == io.EOF {
			//? LOCK THE CONNECTIONS
			csv.connectionsMutex.Lock()
			//? Remove the connection
			csv.connections[firstMessage.Credentials.Username] = nil
			//? Unlock
			csv.connectionsMutex.Unlock()
			break
		}
		if err != nil {
			//? LOCK THE CONNECTIONS
			csv.connectionsMutex.Lock()
			//? Remove the connection
			csv.connections[firstMessage.Credentials.Username] = nil
			//? Unlock
			csv.connectionsMutex.Unlock()
			break
		}
		//? Lock
		csv.usersMutex.RLock()
		csv.connectionsMutex.RLock()
		//? Send the message to all the clients or just for one depending of the receiver
		if newInfo.Receiver != "" {
			for _, user := range csv.users {
				connection := csv.connections[user.Username]
				if user.Username != firstMessage.Credentials.Username && connection != nil && user.Username == newInfo.Receiver {
					(*connection).Send(newInfo)
					break
				}
			}
		} else {
			for _, user := range csv.users {
				connection := csv.connections[user.Username]
				if user.Username != firstMessage.Credentials.Username && connection != nil {
					(*connection).Send(newInfo)
				}
			}
		}
		//? Unlock
		csv.usersMutex.RUnlock()
		csv.connectionsMutex.RUnlock()
	}
	message = &c.Message{}
	message.Msg = fmt.Sprintf("The user %v just left the channel", firstMessage.Credentials.Username)
	message.Sender = "server"
	//? lock
	csv.usersMutex.RLock()
	csv.connectionsMutex.RLock()
	//? Tell everyone a given user left the channel
	for _, user := range csv.users {
		connection := csv.connections[user.Username]
		if connection != nil {
			err := (*connection).Send(message)
			if err != nil {
				fmt.Printf("[Server] Some error occurred %v \n", err)
			}
		}
	}
	//? Unlock
	csv.usersMutex.RUnlock()
	csv.connectionsMutex.RUnlock()
	return nil
}

func main() {
	lis, err := net.Listen("tcp", "localhost:2000")
	if err != nil {
		log.Fatalf("Some error occured listening the server")
	}
	grpc := grpc.NewServer()
	c.RegisterChatServer(
		grpc,
		&ChatServer{users: []*c.UserData{}, connections: make(map[string]*c.Chat_JoinServer)},
	)
	log.Printf("Server listening at %v", lis.Addr())
	if err := grpc.Serve(lis); err != nil {
		log.Fatalf("Some error occured during the lis of the server %v", err)
	}
}
