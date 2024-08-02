package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/carleesmarty/CalcApps/handlers"
	"github.com/carleesmarty/CalcLib"
	"github.com/carleesmarty/CalcLib/Add"
	"github.com/carleesmarty/CalcLib/Divide"
	"github.com/carleesmarty/CalcLib/Mod"
	"github.com/carleesmarty/CalcLib/Multiply"
	"github.com/carleesmarty/CalcLib/Subtract"
)

func main() {

	var calculator CalcLib.Calculator

	fmt.Println("command line args:", os.Args)
	var op string
	flag.StringVar(&op, "op", "", "The mathematical operation to employ")
	fmt.Println("Operation:", op)
	flag.Parse()
	fmt.Println("Operation:", op) // now populated!

	switch {
	case op == "+" || op == "Addition":
		calculator = Add.Addition{}

	case op == "-" || op == "Subtraction":
		calculator = Subtract.Subtraction{}

	case op == "*" || op == "Multiplication":
		calculator = Multiply.Multiplication{}

	case op == "/" || op == "Division":
		calculator = Divide.Division{}

	case op == "%" || op == "Modulo":
		calculator = Mod.Modulo{}
	default:
		log.Fatal("unsupported operation:", op)
	}
	args := flag.Args()
	if len(args) < 2 {
		fmt.Println("Not enough arguments")
		os.Exit(1)
	}

	handler := handlers.NewHandler(calculator, os.Stdout)
	err := handler.Handle(args)
	if err != nil {
		log.Printf("Error handling calculation: %s", err)
		os.Exit(1)
	}

}
