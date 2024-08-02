package handlers

import (
	"fmt"
	"io"
	"strconv"

	"github.com/carleesmarty/CalcLib"
)

type Handler struct {
	calculator CalcLib.Calculator
	out        io.Writer
}

func NewHandler(calculator CalcLib.Calculator, out io.Writer) *Handler {
	return &Handler{out: out, calculator: calculator}
}

func (h *Handler) Handle(args []string) error {

	num1, err1 := strconv.Atoi(args[0])
	num2, err2 := strconv.Atoi(args[1])

	if err1 != nil {
		return fmt.Errorf("invalid arguments: %s", args[0])
	}
	if err2 != nil {
		return fmt.Errorf("invalid arguments: %s", args[1])
	}

	result := h.calculator.Calculate(float64(num1), float64(num2))

	byteResult := []byte(strconv.Itoa(int(result)))
	_, err := h.out.Write(byteResult)
	if err != nil {
		return fmt.Errorf("could not write to output: %w", err)
	}
	return err
}
