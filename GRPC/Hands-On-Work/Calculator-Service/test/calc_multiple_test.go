package test

import (
	c "calc/protofiles"
	"context"
	"testing"
)

func Test_Calc_Multiple(t *testing.T) {
	stream, err := Client.CalcListNumbers(context.Background())
	if err != nil {
		t.Fatalf("error in stream retrieval %v", err)
	}
	for i := 0; i < 10; i++ {
		request := &c.OperationObj{}
		request.Num1 = int32(1)
		request.Num2 = int32(10)
		request.Op = *c.Op_SUMA.Enum()
		err := stream.Send(request)
		if err != nil {
			t.Fatalf("error in sending msg %v", err)
		}
	}
	result, err := stream.CloseAndRecv()
	if err != nil {
		t.Fatalf("something went wrong reading the results %v", err)
	}
	if len(result.ResultList) != 10 {
		t.Fatalf("something is wrong with the length of the list")
	}
	for _, member := range result.ResultList {
		actual := member.Result
		expected := int32(11)
		if actual != expected {
			t.Fatalf("expected %d, we got %d", expected, actual)
		}
	}
}
