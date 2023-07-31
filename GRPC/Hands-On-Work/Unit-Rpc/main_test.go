package testas

import (
	"fmt"
	"os"
	"testing"
)

// Global variable to hold the test data
var testStruct test

type test struct {
	Map map[string]int
}

func TestMain(m *testing.M) {
	// Initialize test data
	testStruct = test{
		Map: make(map[string]int),
	}
	testStruct.Map["test"] = 2
	// Run the tests
	exitCode := m.Run()
	// Exit with the appropriate exit code
	os.Exit(exitCode)
}

func TestValue(t *testing.T) {
	testStruct.Map["test2"] = 3
	fmt.Println(testStruct.Map["test"])
	if 2 != testStruct.Map["test"] {
		t.Errorf("The numbers must be equal")
	}
}

func TestAnother(t *testing.T) {
	if testStruct.Map["test2"] != 3 {
		t.Errorf("It suppose to be equal")
	}
}
