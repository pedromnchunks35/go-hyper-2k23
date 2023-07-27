package main

import (
	"encoding/json"
	fls "filesys/protofiles"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

type filesys struct {
	*fls.UnimplementedFileSharingServer
}

// ? Function to read a certain file and return it back
func (filesys *filesys) GetFile(details *fls.FileDetails, stream fls.FileSharing_GetFileServer) error {
	//? Get the file
	file, err := os.Open("db.json")
	if err != nil {
		return fmt.Errorf("error opening the db: %v", err)
	}
	marshData, err := ioutil.ReadAll(file)
	if err != nil {
		return fmt.Errorf("error reading the file: %v", err)
	}
	//? Unmarshling the file
	var data []File
	err = json.Unmarshal(marshData, &data)
	if err != nil {
		return fmt.Errorf("error decoding the db %v", err)
	}
	name := ""
	//? Loop to find the data
	for _, item := range data {
		if item.Name == details.Name && item.Hash == details.Hash {
			name = item.Name
			break
		}
	}
	if name != "" {
		//? Get the file case the detail got found
		fileData, err := ioutil.ReadFile(fmt.Sprintf("./filestore/%v", details.Name))
		if err != nil {
			return fmt.Errorf("error reading the file from the directory: %v", err)
		}
		//? Send data in batches
		currentBytePos := 0
		tempPos := 0
		i := 1
		for {
			//? STOP POINT
			if currentBytePos >= len(fileData)-1 {
				break
			}
			//? Chunk creation
			tempPos = currentBytePos + int(details.ChunkSize)
			if tempPos > len(fileData)-1 {
				tempPos = len(fileData) - 1
			}
			chunk := fileData[currentBytePos:tempPos]
			currentBytePos = tempPos
			//? Message creation
			chunkToSend := fls.FileContent{}
			chunkToSend.ChunkNumber = int32(i)
			i++
			chunkToSend.DefaultChunkSize = int32(details.ChunkSize)
			chunkToSend.ChunkSize = int32(len(chunk))
			chunkToSend.Data = chunk
			chunkToSend.Hash = details.Hash
			chunkToSend.Name = details.Name
			//? Send the message
			err := stream.Send(&chunkToSend)
			if err != nil {
				return fmt.Errorf("something went wrong sending the chunk")
			}
		}
		return nil
	} else {
		return fmt.Errorf("file not found")
	}
}

// ? Function that will save files
func (filesys *filesys) SaveFile(stream fls.FileSharing_SaveFileServer) error {
	//? Init the full byte array
	result := []byte{}
	//? Get the first details
	firstMessage, err := stream.Recv()
	if err != nil {
		return fmt.Errorf("there is nothing on the request: %v", err)
	}
	result = append(result, firstMessage.Data...)
	//? Make the normal iteraction
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("something went wrong on reading the message: %v", err)
		}
		result = append(result, msg.Data...)
	}
	//? Store the file
	err = ioutil.WriteFile(fmt.Sprintf("./filestore/%v", firstMessage.Name), result, 0644)
	if err != nil {
		return fmt.Errorf("something went wrong saving the file: %v", err)
	}
	//? Save the file details in the db
	file, err := os.Open("db.json")
	if err != nil {
		return fmt.Errorf("error opening the database %v", err)
	}
	marshData, err := ioutil.ReadAll(file)
	if err != nil {
		return fmt.Errorf("error reading the database %v", err)
	}
	var data []File
	err = json.Unmarshal(marshData, &data)
	newDetail := File{Name: firstMessage.Name, Hash: firstMessage.Hash}
	if err != nil {
		return fmt.Errorf("error unmarshling the file %v", err)
	}
	data = append(data, newDetail)
	//? Marshall again with intend for formatting the json
	newData, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return fmt.Errorf("erro encoding the data back to json %v", err)
	}
	//? Write the file
	err = ioutil.WriteFile("db.json", newData, 0644)
	if err != nil {
		return fmt.Errorf("error making the write of the file")
	}
	//? Send confirmation
	confirmation := &fls.Confirmation{}
	confirmation.Msg = "File stored sucessfully"
	confirmation.Path = fmt.Sprintf("./filestore/%v", firstMessage.Name)
	confirmation.Status = 200
	stream.SendAndClose(confirmation)
	return nil
}

type File struct {
	Name string `json:"name"`
	Hash string `json:"hash"`
}

var PORT = flag.Int("port", 2000, "The port of the server")

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%v", *PORT))
	if err != nil {
		log.Fatalf("Something went wrong when making the server accessible %v", err)
	}
	grpc := grpc.NewServer()
	fls.RegisterFileSharingServer(grpc, &filesys{})
	log.Printf("Server starting at %v", lis.Addr())
	err = grpc.Serve(lis)
	if err != nil {
		log.Fatalf("Something went wrong making the server available: %v", err)
	}
}
