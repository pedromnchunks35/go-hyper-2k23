package main

import (
	pf "calc/protofiles"
	"context"
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type operationObjCmd struct {
	num1 int32
	num2 int32
	op   string
}

type listToCalc []*operationObjCmd

// ? String version of our array
func (list *listToCalc) String() string {
	s := ""
	for _, op := range *list {
		s += fmt.Sprintf("{Numbers:[%v,%v],Operation:%v}", op.num1, op.num2, op.op)
	}
	return s
}

// ? Creation of our list object
func (list *listToCalc) Set(value string) error {
	items := strings.Split(value, ",")
	if len(items) != 3 {
		log.Fatalf("You need to provide 3 items")
	}
	item := &operationObjCmd{}
	num1, _ := strconv.Atoi(items[0])
	num2, _ := strconv.Atoi(items[1])
	item.num1 = int32(num1)
	item.num2 = int32(num2)
	item.op = items[2]
	*list = append(*list, item)
	return nil
}

// ? Creation of a single object
func (single operationObjCmd) String() string {
	return fmt.Sprintf("{Numbers:[%v,%v],Operation:%v}", single.num1, single.num2, single.op)
}

// ? Single object set of value in flags
func (item *operationObjCmd) Set(value string) error {
	items := strings.Split(value, ",")
	if len(items) != 3 {
		log.Fatalf("You need to provide 3 items")
	}
	num1, _ := strconv.Atoi(items[0])
	num2, _ := strconv.Atoi(items[1])
	item.num1 = int32(num1)
	item.num2 = int32(num2)
	item.op = items[2]
	return nil
}

var (
	serverAddr      = flag.String("addr", "localhost:50051", "The server we wish to communicate with url")
	numsToCalculate listToCalc
	toCalc          operationObjCmd
)

func getEnum(value string) pf.Op {
	switch value {
	case pf.Op_DIVI.String():
		return pf.Op_DIVI
	case pf.Op_MULT.String():
		return pf.Op_MULT
	case pf.Op_SUBT.String():
		return pf.Op_SUBT
	case pf.Op_SUMA.String():
		return pf.Op_SUMA
	default:
		return pf.Op_SUMA
	}
}

// ? Calculate without list
func simpleCalc(client pf.DoMathClient, operation *pf.OperationObj) {
	log.Printf("Getting the calculation response for %v", toCalc.String())
	//? Create some context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//? release resources
	defer cancel()
	//? Get the result
	result, err := client.CalcNumbers(ctx, operation)
	//? Handle error
	if err != nil {
		log.Fatalf("Some error occured in a simpleCalc %v", err)
	}
	log.Printf("The result from this call is: %v\n", result.Result)
}

// ? Calculate with list
func multipleCalc(client pf.DoMathClient, operations listToCalc) {
	log.Printf("Getting the calculation response for %v", numsToCalculate.String())
	//? create some context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//? Get a stream
	stream, err := client.CalcListNumbers(ctx)
	if err != nil {
		log.Fatalf("Error getting the stream for multiple calc: %v", err)
	}
	//? Conceive the request
	for _, opt := range operations {
		newOperation := &pf.OperationObj{}
		newOperation.Num1 = opt.num1
		newOperation.Num2 = opt.num2
		newOperation.Op = getEnum(opt.op)
		stream.Send(newOperation)
	}
	//? Close and receive the response
	result, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Something went wrong with the request: %v", err)
	}
	fmt.Printf("The result is: %v", result)
}

func main() {
	//? handle the flags
	flag.Var(&numsToCalculate, "operations", "operations flag")
	flag.Var(&toCalc, "operation", "flag for only one operation")
	flag.Parse()
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	//? Create the connection
	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("Fail to dial: %v", err)
	}
	defer conn.Close()
	//? create the client
	client := pf.NewDoMathClient(conn)
	//? For multicalc
	if len(numsToCalculate) != 0 {
		multipleCalc(client, numsToCalculate)
	}
	//? For single calc
	if toCalc.op != "" {
		operation := &pf.OperationObj{}
		operation.Num1 = toCalc.num1
		operation.Num2 = toCalc.num2
		operation.Op = getEnum(toCalc.op)
		simpleCalc(client, operation)
	}
}
