package test

import (
	c "chat/protofiles"
	"context"
	"fmt"
	"io"
	"strings"
	"testing"
)

func Test_Register(t *testing.T) {
	userData := &c.UserData{}
	userData.Username = "teste"
	userData.Password = "12341234"
	result, err := Client.Register(context.Background(), userData)
	if err != nil {
		t.Fatalf("it should not throw a error, the register is valid %v", err)
	}
	if !strings.Contains(result.Msg, "The user teste got registered") {
		t.Fatalf("it should throw a success message")
	}
}

func Test_Register_Invalid(t *testing.T) {
	userData := &c.UserData{}
	userData.Username = "teste"
	userData.Password = ""
	_, err := Client.Register(context.Background(), userData)
	if err == nil || !strings.Contains(err.Error(), "the user is already registered") {
		t.Fatalf("it should throw a error with a message saying that it is already registered")
	}
}

func Test_Join_Invalid(t *testing.T) {
	//? join the first user with wrong credentials first
	streamTeste1, err := Client.Join(context.Background())
	if err != nil {
		t.Fatalf("error in getting the stream for teste1 %v", err)
	}
	msg := &c.Message{}
	msg.Credentials = &c.UserData{Username: "teste", Password: "123"}
	err = streamTeste1.Send(msg)
	if err != nil {
		t.Fatalf("error sending the credentials %v", err)
	}
	_, err = streamTeste1.Recv()
	if !strings.Contains(err.Error(), "please provide valid credentials") {
		t.Fatalf("it should return a error on the first message because of bad credentials")
	}
	err = streamTeste1.Send(msg)
	if err != io.EOF {
		t.Fatalf("it should throw a error because the credentials were wrong, so the connection got closed")
	}
}

func Test_Join(t *testing.T) {
	//? Add one more user
	userData := &c.UserData{}
	userData.Username = "teste2"
	userData.Password = "12341234"
	result, err := Client.Register(context.Background(), userData)
	if err != nil {
		t.Fatalf("it should not throw a error, the register is valid %v", err)
	}
	if !strings.Contains(result.Msg, "The user teste2 got registered") {
		t.Fatalf("it should throw a success message")
	}
	//? join the first user
	streamTeste1, err := Client.Join(context.Background())
	if err != nil {
		t.Fatalf("error in getting the stream for teste1 %v", err)
	}
	msg := &c.Message{}
	msg.Credentials = &c.UserData{Username: "teste", Password: "12341234"}
	err = streamTeste1.Send(msg)
	if err != nil {
		t.Fatalf("error sending the credentials %v", err)
	}
	//? Join the secound user
	msg.Credentials = userData
	streamTeste2, err := Client.Join(context.Background())
	if err != nil {
		t.Fatalf("error in getting the stream for teste2 %v", err)
	}
	//? The first user should receive a message saying that he joined the channel
	res, err := streamTeste1.Recv()
	if err != nil {
		t.Fatalf("should not throw a error at this state, teste joining the channel %v", err)
	}
	if !strings.Contains(res.Msg, fmt.Sprintf("The user %v just joined the channel", "teste")) {
		t.Fatalf("should receive a message saying that user2 joined the channel")
	}
	//? user2 join the channel
	msg = &c.Message{}
	msg.Credentials = &c.UserData{Username: "teste2", Password: "12341234"}
	err = streamTeste2.Send(msg)
	if err != nil {
		t.Fatalf("error sending the credentials %v", err)
	}
	//? The first user should receive a message saying that the secound user joined the channel
	res, err = streamTeste1.Recv()
	if err != nil {
		t.Fatalf("should not throw a error at this state, teste2 joining the channel %v", err)
	}
	if !strings.Contains(res.Msg, fmt.Sprintf("The user %v just joined the channel", userData.Username)) {
		t.Fatalf("should receive a message saying that user2 joined the channel")
	}
	//? Secound user should receive the message that himself joined the channel
	//? The first user should receive a message saying that he joined the channel
	res, err = streamTeste2.Recv()
	if err != nil {
		t.Fatalf("should not throw a error at this state, teste2 joining the channel %v", err)
	}
	if !strings.Contains(res.Msg, fmt.Sprintf("The user %v just joined the channel", "teste2")) {
		t.Fatalf("should receive a message saying that user2 joined the channel")
	}
	//? User2 sends a message and the user1 receives the message in a global way
	normalMessage := &c.Message{}
	normalMessage.Msg = "Hello my dear"
	err = streamTeste2.Send(normalMessage)
	if err != nil {
		t.Fatalf("should not return a error when we send a normal message from user2 to user1 %v", err)
	}
	resToNormalMessage, err := streamTeste1.Recv()
	if err != nil {
		t.Fatalf("should not return a error when receiving the normal message %v", err)
	}
	if !strings.Contains(resToNormalMessage.Msg, "Hello my dear") {
		t.Fatalf("it hould receive the normal message from the user2, this is a global message")
	}
	//? Joining a third user
	//? Add one more user
	userData = &c.UserData{}
	userData.Username = "teste3"
	userData.Password = "12341234"
	result, err = Client.Register(context.Background(), userData)
	if err != nil {
		t.Fatalf("it should not throw a error, the register is valid %v", err)
	}
	if !strings.Contains(result.Msg, "The user teste3 got registered") {
		t.Fatalf("it should throw a success message")
	}
	//? Join the third user
	streamTeste3, err := Client.Join(context.Background())
	if err != nil {
		t.Fatalf("error in getting the stream for teste3 %v", err)
	}
	msg = &c.Message{}
	msg.Credentials = &c.UserData{Username: "teste3", Password: "12341234"}
	err = streamTeste3.Send(msg)
	if err != nil {
		t.Fatalf("error sending the credentials %v", err)
	}
	//? User2 sends a message and the user1 receives it privately
	normalMessage.Receiver = "teste"
	err = streamTeste2.Send(normalMessage)
	if err != nil {
		t.Fatalf("error when sending a private message to the user1 %v", err)
	}
	//? It should receive the private message from the user2
	res, err = streamTeste1.Recv()
	if err != nil {
		t.Fatalf("error when reading the message of the new normal private message from user2 to user1 %v", err)
	}
	if !strings.Contains(res.Msg, "Hello my dear") || res.Receiver != "teste" {
		t.Fatalf("it should receive the message and the receiver must be the user1")
	}
	//? User1 must read the message that the third user joined the server
	res, err = streamTeste1.Recv()
	if err != nil {
		t.Fatalf("error when reading the message of the third user join in stream1 %v", err)
	}

	if !strings.Contains(res.Msg, fmt.Sprintf("The user %v just joined the channel", "teste3")) {
		t.Fatalf("it should receive the message saying that the third user joined the chat")
	}
	//? The third user cannot receive the message,lets send a message global again to see that
	//? User2 sends a message and the user1 receives the message in a global way
	normalMessage2 := &c.Message{}
	normalMessage2.Msg = "Third user will read this message first because the last one was private to user1"
	err = streamTeste2.Send(normalMessage2)
	if err != nil {
		t.Fatalf("should not return a error when we send a normal message to everyone %v", err)
	}
	//? Third user receives the message that he joined the server
	res, err = streamTeste3.Recv()
	if err != nil {
		t.Fatalf("should not throw a error in the receive of the global message %v", err)
	}
	if !strings.Contains(res.Msg, fmt.Sprintf("The user %v just joined the channel", "teste3")) {
		t.Fatalf("third user must receive the global message saying he joined the channel")
	}
	//? Third user receives the global message
	res, err = streamTeste3.Recv()
	if err != nil {
		t.Fatalf("should not throw a error in the receive of the global message %v", err)
	}
	if !strings.Contains(res.Msg, "Third user will read this message first because the last one was private to user1") {
		t.Fatalf("third user must receive the global message isntead of the private one")
	}
}
