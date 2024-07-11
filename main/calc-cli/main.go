package main

import (
	"log"
	"os"

	"github.com/carleesmarty/CalcLib"
)

func main() {

	//firstInt, err1 := strconv.Atoi(os.Args[1])
	//secondInt, err2 := strconv.Atoi(os.Args[2])
	//fmt.Println(os.Args[0], os.Args[1])
	//if err1 != nil || err2 != nil {
	//	panic("Not a valid int.")
	//}
	//fmt.Println(CalcLib.Addition{}.Calculate(firstInt, secondInt))
	//
	args := []string{"add", "10", "67"}
	calculator := CalcLib.Addition{}

	handler := NewHandler(calculator, os.Stdout)
	err := handler.Handle(args)
	if err != nil {
		log.Printf("Error handling calculation: %s", err)
		os.Exit(1)
	}

}
