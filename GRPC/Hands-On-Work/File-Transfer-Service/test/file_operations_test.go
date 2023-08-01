package test

import (
	"bytes"
	"context"
	"encoding/json"
	fls "filesys/protofiles"
	"io"
	"os"
	"testing"
)

type File struct {
	Name string `json:"name"`
	Hash string `json:"hash"`
}

func Test_File_Save(t *testing.T) {
	stream, err := Client.SaveFile(context.Background())
	if err != nil {
		t.Fatalf("some error occured trying to get the stream %v", err)
	}
	file, err := os.ReadFile("./assets/image.jpg")
	if err != nil {
		t.Fatalf("error reading the image %v", err)
	}
	currentPos := 0
	tempPos := 0
	i := 1
	//? Size in bytes
	fileChunkSize := 30
	for {
		if currentPos == len(file) {
			break
		}
		if tempPos+fileChunkSize >= len(file) {
			tempPos = len(file)
		} else {
			tempPos += fileChunkSize
		}
		msg := &fls.FileContent{}
		msg.ChunkNumber = int32(i)
		msg.ChunkSize = int32(fileChunkSize)
		msg.DefaultChunkSize = int32(fileChunkSize)
		msg.Name = "image.jpg"
		msg.Data = file[currentPos:tempPos]
		currentPos = tempPos
		i++
		err := stream.Send(msg)
		if err != nil {
			t.Fatalf("something went wrong sending the file content %v", err)
		}
	}
	response, err := stream.CloseAndRecv()
	if err != nil {
		t.Fatalf("something went wrong with the response %v", err)
	}
	if response.Msg != "File stored sucessfully" {
		t.Fatalf("must answer with sucessfull message")
	}
	if response.Path != "./filestore/image.jpg" {
		t.Fatalf("must answer with the correct path")
	}
	if response.Status != int32(200) {
		t.Fatalf("must answer with the correct status")
	}
	storedFile, err := os.ReadFile("./filestore/image.jpg")
	if err != nil {
		t.Fatalf("there must be a file in it %v", err)
	}
	if !bytes.Equal(file, storedFile) {
		t.Fatalf("the files must be the same")
	}
	storedJson, err := os.ReadFile("./db.json")
	if err != nil {
		t.Fatalf("something went wrong reading the db %v", err)
	}
	var content []File
	err = json.Unmarshal(storedJson, &content)
	if err != nil {
		t.Fatalf("something went wrong unmarshling the data")
	}
	found := false
	for _, cont := range content {
		if cont.Name == "image.jpg" && cont.Hash == "" {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("it have the name of the file in the db and the hash as well")
	}
}

func Test_Get(t *testing.T) {
	fileDetails := &fls.FileDetails{}
	fileDetails.ChunkSize = 20
	fileDetails.Hash = ""
	fileDetails.Name = "image.jpg"
	stream, err := Client.GetFile(context.Background(), fileDetails)
	if err != nil {
		t.Fatalf("error during the getfile steam %v", err)
	}
	result := []byte{}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		result = append(result, res.Data...)
	}
	file, err := os.ReadFile("./assets/image.jpg")
	if err != nil {
		t.Fatalf("error reading the image %v", err)
	}
	if !bytes.Equal(result, file) {
		t.Fatalf("the file is not the same %v", err)
	}
	//? Clear the project
	os.Remove("./filestore/image.jpg")
	content := []File{}
	result, err = json.Marshal(content)
	if err != nil {
		t.Fatalf("clear of the db went wrong")
	}
	os.WriteFile("./db.json", result, 0644)
}
