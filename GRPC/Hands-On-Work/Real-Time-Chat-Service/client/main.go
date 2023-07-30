package main

import (
	"bufio"
	c "chat/protofiles"
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// ? Function to simply register a given user
func Register(client c.ChatClient, userData *c.UserData) error {
	ctx := context.Background()
	conf, err := client.Register(ctx, userData)
	if err != nil {
		return fmt.Errorf("some error occured registering the user: %v", err)
	}
	log.Printf("[Server] %v", conf.Msg)
	return nil
}

// ? Function to receive the messages
func receiveMessages(stream *c.Chat_JoinClient, errCh chan<- error) error {
	for {
		message, err := (*stream).Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			errCh <- err
			return fmt.Errorf("some error occured during the message reading: %v", err)
		}
		if message.Receiver != "" {
			fmt.Printf("[%v -> Me (private msg)] %v \n", message.Sender, message.Msg)
		} else {
			fmt.Printf("[%v] %v\n", message.Sender, message.Msg)
		}
	}
	return nil
}

// ? Function to join the channel
func Join(client c.ChatClient, firstMessage *c.Message) error {
	ctx := context.Background()
	stream, err := client.Join(ctx)
	stream.Send(firstMessage)
	if err != nil {
		return fmt.Errorf("some error occured in the login: %v", err)
	} else {
		//? Read the confirmation
		confirmation, err := stream.Recv()
		if err != nil {
			return fmt.Errorf("some error occured in the login: %v", err)
		}
		fmt.Println("*******************************************************************************************************")
		fmt.Printf("[Server] %v\n", confirmation.Msg)
		fmt.Println("Please write /tell <name of the user> <message>, in case you want to send something to a restrict user")
		fmt.Println("Please write /exit, in case you desire to exit the chat")
		fmt.Println("*******************************************************************************************************")
		//? Create a channel
		errCh := make(chan error)
		//? Create a thread for reading messages
		go receiveMessages(&stream, errCh)
		//? Listen to errors
		go func() {
			for err := range errCh {
				if err != nil {
					fmt.Println("[ERROR IN CHANNEL MSG RECEIVAL] ", err)
				}
			}
		}()
		//? Create a new bufio.Reader to read from standard input (console)
		reader := bufio.NewReader(os.Stdin)
		for {
			//? Get the message
			message, err := reader.ReadString('\n')
			message = strings.TrimRight(message, "\r\n")
			message = strings.TrimRight(message, "\n")
			if err != nil {
				return fmt.Errorf("unexpected error just occured %v", err)
			}
			//? Exit part
			if message == "/exit" {
				break
			}
			//? Check if there is any commmand
			command := strings.Split(message, " ")
			if len(command) >= 3 && command[0] == "/tell" && command[1] != "" && command[2] != "" {
				//? Send the message to a particular person
				msg := &c.Message{}
				msg.Msg = ""
				for i := 2; i < len(command); i++ {
					msg.Msg += fmt.Sprintf(" %v", command[i])
				}
				msg.Sender = firstMessage.Credentials.Username
				msg.Receiver = command[1]
				err := stream.Send(msg)
				//? Print error or the msg we just sended
				if err != nil {
					fmt.Printf("[Error] Some error occured trying to send the message: %v \n", err)
				} else {
					fmt.Printf("[Me -> %v (private msg)] %v\n", command[1], msg.Msg)
				}
			} else {
				//? Case there is no command broadcast for all
				msg := &c.Message{}
				msg.Msg = message
				msg.Sender = firstMessage.Credentials.Username
				err := stream.Send(msg)
				//? Print error or the msg we just sended
				if err != nil {
					fmt.Printf("[Error] Some error occured trying to send the message: %v \n", err)
				} else {
					fmt.Printf("[Me] %v\n", message)
				}
			}
		}
	}
	fmt.Println("*******************************************************************************************************")
	fmt.Println("Logging out.. thanks for chating with us :)")
	fmt.Println("*******************************************************************************************************")
	//? Close the stream and left
	err = stream.CloseSend()
	if err != nil {
		return fmt.Errorf("error closing the stream %v", err)
	}
	return nil
}

var addr = flag.String("addr", "localhost", "The addr")

func main() {
	flag.Parse()
	//? initiate the conn
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial(fmt.Sprintf("%v:2000", *addr), opts...)
	if err != nil {
		log.Fatalf("some error occured making the dial %v", err)
	}
	defer conn.Close()
	client := c.NewChatClient(conn)
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("*******************************************************************************************************")
		fmt.Println("1- Register")
		fmt.Println("2- Join")
		fmt.Println("3- Exit")
		fmt.Println("*******************************************************************************************************")
		opt, err := reader.ReadString('\n')
		opt = strings.TrimRight(opt, "\r\n")
		opt = strings.TrimRight(opt, "\n")
		if err != nil {
			log.Fatalf("Some error occured reading the opt: %v", err)
		}
		//? Case it is a register
		if opt == "1" {
			//? Grab register details and execute the func of register
			cred := &c.UserData{}
			fmt.Println("*******************************************************************************************************")
			fmt.Println("Please write down the username")
			fmt.Println("*******************************************************************************************************")
			username, err := reader.ReadString('\n')
			username = strings.TrimRight(username, "\r\n")
			username = strings.TrimRight(username, "\n")
			if err != nil {
				log.Fatalf("Some error occured reading the username: %v", err)
			} else {
				cred.Username = username
				fmt.Println("*******************************************************************************************************")
				fmt.Println("Please write down the password")
				fmt.Println("*******************************************************************************************************")
				password, err := reader.ReadString('\n')
				password = strings.TrimRight(password, "\r\n")
				password = strings.TrimRight(password, "\n")
				if err != nil {
					log.Fatalf("Some error occured reading the password: %v", err)
				} else {
					cred.Password = password
					err := Register(client, cred)
					if err != nil {
						log.Fatalf("Some error occured with the register: %v", err)
					}
				}
			}
		} else if opt == "2" {
			//? Grab register details and execute the func of register
			cred := &c.UserData{}
			fmt.Println("*******************************************************************************************************")
			fmt.Println("Please write down the username")
			fmt.Println("*******************************************************************************************************")
			username, err := reader.ReadString('\n')
			username = strings.TrimRight(username, "\r\n")
			username = strings.TrimRight(username, "\n")
			if err != nil {
				log.Fatalf("Some error occured reading the username: %v", err)
			} else {
				cred.Username = username
				fmt.Println("*******************************************************************************************************")
				fmt.Println("Please write down the password")
				fmt.Println("*******************************************************************************************************")
				password, err := reader.ReadString('\n')
				password = strings.TrimRight(password, "\r\n")
				password = strings.TrimRight(password, "\n")
				if err != nil {
					log.Fatalf("Some error occured reading the password: %v", err)
				} else {
					cred.Password = password
					//? Put the credentials inside of the first message
					firstMessage := &c.Message{}
					firstMessage.Credentials = cred
					//? Make here the join
					err = Join(client, firstMessage)
					if err != nil {
						log.Fatalf("%v", err)
					}
				}
			}
		} else if opt == "3" {
			break
		}
	}

}
