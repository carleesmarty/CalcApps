package handlers

import (
	"bytes"
	"errors"
	"reflect"
	"testing"
)

func TestHandlerHappyPath(t *testing.T) {
	calculator := &fakeCalculator{output: 42}
	buffer := &bytes.Buffer{}
	handler := NewHandler(calculator, buffer)

	err := handler.Handle([]string{"", "1", "4"})

	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(calculator.inputs, []int{1, 4}) {
		t.Errorf("Incorrect inputs: %v", calculator.inputs)
	}
	if buffer.String() != "42" {
		t.Errorf("Incorrect outputs: %s", buffer.String())
	}
}

func TestHandlerArgumentsLengthFail(t *testing.T) {

	calculator := &fakeCalculator{output: 42}
	buffer := &bytes.Buffer{}
	handler := NewHandler(calculator, buffer)

	err := handler.Handle([]string{"", "1"})
	if err == nil {
		t.Errorf("String contains too few argument")
	}
}

func TestHandlerArgumentsLengthPass(t *testing.T) {
	calculator := &fakeCalculator{output: 42}
	buffer := &bytes.Buffer{}
	handler := NewHandler(calculator, buffer)

	err := handler.Handle([]string{"", "1", "4"})
	if err != nil {
		t.Error(err)
	}
}

func TestHandlerArgumentsTypeFail(t *testing.T) {
	calculator := &fakeCalculator{output: 42}
	buffer := &bytes.Buffer{}
	handler := NewHandler(calculator, buffer)

	err := handler.Handle([]string{"", "1", "abcd"})
	if err == nil {
		t.Errorf("Incorrect inputs: integers needed")
	}
}

func TestHandlerArgumentsTypeFail2(t *testing.T) {
	calculator := &fakeCalculator{output: 42}
	buffer := &bytes.Buffer{}
	handler := NewHandler(calculator, buffer)

	err := handler.Handle([]string{"", "abcd", "3"})
	if err == nil {
		t.Errorf("Incorrect inputs: integers needed")
	}
}

func TestHandleWriteError(t *testing.T) {
	errorWriter := &ErrorWriter{}
	calculator := &fakeCalculator{output: 42}
	handler := NewHandler(calculator, errorWriter)
	err := handler.Handle([]string{"add", "4", "1"})

	if err == nil {
		t.Errorf("Expected and error.")
	}
	if !errors.Is(err, writeError) {
		t.Errorf("Expected writeError, got: %v", err)
	}

}

type fakeCalculator struct {
	inputs []int
	output int
}

func (f *fakeCalculator) Calculate(a, b int) int {
	f.inputs = append(f.inputs, a, b)
	return f.output
}

// ErrorWriter is a custom writer that always returns an error
type ErrorWriter struct{}

func (e *ErrorWriter) Write(p []byte) (n int, err error) {
	return 0, writeError
}

var writeError = errors.New("write error")
