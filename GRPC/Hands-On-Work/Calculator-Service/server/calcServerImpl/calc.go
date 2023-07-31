package calcServerImpl

import (
	pf "calc/protofiles"
	"context"
	"flag"
	"fmt"
	"io"
)

// ? The main struct for the operations
type CalculationServer struct {
	pf.UnimplementedDoMathServer
}

var PORT = flag.Int("port", 50051, "The port of the calculation server")

// ? Error structure
type MyError string

func (e MyError) Error() string {
	return string(e)
}

// ? The function that does the calculation
func calculate(
	num1 int32,
	num2 int32,
	operation pf.Op,
) (int32, error) {
	switch operation.String() {
	case "MULT":
		return (num1 * num2), nil
	case "SUBT":
		return (num1 - num2), nil
	case "SUMA":
		return (num1 + num2), nil
	case "DIVI":
		return (num1 / num2), nil
	default:
		return 0, MyError("The option must be a valid one")
	}
}

// ? Makes calculation with a given operation
func (calc *CalculationServer) CalcNumbers(ctx context.Context, operation *pf.OperationObj) (*pf.ResultObj, error) {
	//? Make the calculation
	calcResult, err := calculate(
		operation.Num1,
		operation.Num2,
		operation.Op,
	)
	//? Error handling
	if err != nil {
		return nil, err
	}
	result := &pf.ResultObj{}
	result.Result = calcResult
	result.Msg = "The calculation is done"
	return result, nil
}

// ? Makes calculation of a stream of numbers
func (calc *CalculationServer) CalcListNumbers(stream pf.DoMath_CalcListNumbersServer) error {
	//? init the list
	list := &pf.ResultObjList{}
	list.ResultList = []*pf.ResultObj{}
	//? Loop all over the streams to get operations
	for {
		//? Get single opt
		operationObj, err := stream.Recv()
		//? Case we reach the end successfully
		if err == io.EOF {
			break
		}
		//? Error handling
		if err != nil {
			return MyError(fmt.Sprintf("Something went wrong with our operation: %v", err))
		}
		//? Calculation
		singleOperationResult, err := calculate(
			operationObj.Num1,
			operationObj.Num2,
			operationObj.Op,
		)
		//? Error handling
		if err != nil {
			return MyError(fmt.Sprintf("Something went wrong with our operation: %v", err))
		}
		//? Assign result
		singleResult := &pf.ResultObj{}
		singleResult.Result = singleOperationResult
		singleResult.Msg = "Operation done"
		list.ResultList = append(list.ResultList, singleResult)
	}
	stream.SendAndClose(list)
	return nil
}
