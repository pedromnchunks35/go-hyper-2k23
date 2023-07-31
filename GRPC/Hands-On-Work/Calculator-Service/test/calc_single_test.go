package test

import (
	c "calc/protofiles"
	"context"
	"testing"
)

func Test_Calc_SUM(t *testing.T) {
	request := &c.OperationObj{}
	request.Num1 = int32(1)
	request.Num2 = int32(10)
	request.Op = *c.Op_SUMA.Enum()
	result, err := Client.CalcNumbers(context.Background(), request)
	if err != nil {
		t.Fatalf("the function call failed for sum %v", err)
	}
	if result.Result != int32(11) {
		t.Errorf("expected %d, got %d", int32(11), result.Result)
	}
}

func Test_Calc_MULT(t *testing.T) {
	request := &c.OperationObj{}
	request.Num1 = int32(1)
	request.Num2 = int32(10)
	request.Op = *c.Op_MULT.Enum()
	result, err := Client.CalcNumbers(context.Background(), request)
	if err != nil {
		t.Fatalf("the function call failed for sum %v", err)
	}
	actual := result.Result
	expected := int32(10)
	if actual != expected {
		t.Errorf("expected %d, got %d", expected, actual)
	}
}

func Test_Calc_SUBT(t *testing.T) {
	request := &c.OperationObj{}
	request.Num1 = int32(1)
	request.Num2 = int32(10)
	request.Op = *c.Op_SUBT.Enum()
	result, err := Client.CalcNumbers(context.Background(), request)
	if err != nil {
		t.Fatalf("the function call failed for sum %v", err)
	}
	actual := result.Result
	expected := int32(-9)
	if actual != expected {
		t.Errorf("expected %d, got %d", expected, actual)
	}
}

func Test_Calc_DIVI(t *testing.T) {
	request := &c.OperationObj{}
	request.Num1 = int32(6)
	request.Num2 = int32(3)
	request.Op = *c.Op_DIVI.Enum()
	result, err := Client.CalcNumbers(context.Background(), request)
	if err != nil {
		t.Fatalf("the function call failed for sum %v", err)
	}
	actual := result.Result
	expected := int32(2)
	if actual != expected {
		t.Errorf("expected %d, got %d", expected, actual)
	}
}
