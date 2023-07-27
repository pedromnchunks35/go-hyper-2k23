package main

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	fls "filesys/protofiles"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func SaveFile(client fls.FileSharingClient) error {
	//? GET THE FILE
	file, err := ioutil.ReadFile(*path)
	if err != nil {
		return fmt.Errorf("something went wrong when reading up the file")
	}
	if *chunkSize == 0 {
		return fmt.Errorf("you need to provide a correct chunkSize")
	}
	if *fileName == "" {
		return fmt.Errorf("please give a valid name to the file")
	}
	ctx := context.Background()
	stream, err := client.SaveFile(ctx)
	if err != nil {
		return fmt.Errorf("something went wrong creating the connection: %v", err)
	}
	hasher := sha256.New()
	hasher.Write(file)
	hashBytes := hasher.Sum(nil)
	newHash := hex.EncodeToString(hashBytes)
	log.Printf("Storing a file with hash: %v", newHash)
	//? CREATE THE CHUNKS
	currentPos := 0
	tempPos := 0
	i := 1
	for {
		//? CHUNK CREATION
		if currentPos >= len(file)-1 {
			break
		}
		tempPos = currentPos + *chunkSize
		if tempPos > len(file)-1 {
			tempPos = len(file) - 1
		}
		chunk := file[currentPos:tempPos]
		currentPos = tempPos
		//? REQUEST CREATION
		fileContentToSend := &fls.FileContent{}
		fileContentToSend.ChunkNumber = int32(i)
		i++
		fileContentToSend.ChunkSize = int32(len(chunk))
		fileContentToSend.DefaultChunkSize = int32(*chunkSize)
		fileContentToSend.Data = chunk
		fileContentToSend.Name = *fileName
		fileContentToSend.Hash = newHash
		stream.Send(fileContentToSend)
	}
	//? Get the reply
	reply, err := stream.CloseAndRecv()
	if err != nil {
		return fmt.Errorf("something went wrong with the response receival: %v", err)
	}
	log.Printf("%v\n", reply)
	return nil
}

func GetFile(client fls.FileSharingClient, details *fls.FileDetails) error {
	//? MAKE THE CONNECTION
	ctx := context.Background()
	stream, err := client.GetFile(ctx, details)
	if err != nil {
		return fmt.Errorf("something went wrong with the connection: %v", err)
	}
	//? Init the file retrieval
	result := []byte{}
	//? Get the first message
	firstMessage, err := stream.Recv()
	if err != nil {
		return fmt.Errorf("there is no file: %v", err)
	}
	result = append(result, firstMessage.Data...)
	//? Get the other chunks
	for {
		chunk, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("something went wrong reading the message: %v", err)
		}
		result = append(result, chunk.Data...)
	}
	//? Write the file
	err = ioutil.WriteFile(fmt.Sprintf("%v%v", *path, details.Name), result, 0644)
	if err != nil {
		return fmt.Errorf("some error occured during file storing: %v", err)
	}
	err = stream.CloseSend()
	if err != nil {
		return fmt.Errorf("error closing the stream: %v", err)
	}
	log.Printf("Success retrieving the file :)")
	return nil
}

var (
	path      = flag.String("path", "", "relative path to the file we want to send")
	chunkSize = flag.Int("chunk", 0, "the size of the batch")
	hash      = flag.String("hash", "", "hash to get the file")
	fileName  = flag.String("name", "", "name of the file")
	storefile = flag.Bool("store", false, "if we want to store a file")
	getfile   = flag.Bool("get", false, "if we want to get a file")
)

func main() {
	flag.Parse()
	opts := []grpc.DialOption{}
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial("localhost:2000", opts...)
	if err != nil {
		log.Fatalf("something went wrong when making the dial: %v", err)
	}
	defer conn.Close()
	client := fls.NewFileSharingClient(conn)
	if *storefile {
		if *path == "" || *chunkSize == 0 || *fileName == "" {
			log.Fatalf("you need to set up path,chunkSize and name correctly")
		}
		err := SaveFile(client)
		if err != nil {
			log.Fatalf("Some error occurred at the final of the procedure: %v", err)
		}
	} else if *getfile {
		if *path == "" || *chunkSize == 0 || *fileName == "" || *hash == "" {
			log.Fatalf("you need to set up path,chunkSize,hash and name correctly")
		}
		fileDetails := &fls.FileDetails{}
		fileDetails.ChunkSize = int32(*chunkSize)
		fileDetails.Hash = *hash
		fileDetails.Name = *fileName
		err := GetFile(client, fileDetails)
		if err != nil {
			log.Fatalf("Some error occurred at the final of the procedure: %v", err)
		}
	}
}
