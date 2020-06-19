package main

import (
	"fmt"
	"github.com/chino976/go-meli-integration/TP_03_Hello_World"
	"github.com/chino976/go-meli-integration/TP_04_Calc"
	"github.com/chino976/go-meli-integration/TP_05_Ping"
	"github.com/chino976/go-meli-integration/TP_06_07_Calc_Mod_Test"
	TP_08 "github.com/chino976/go-meli-integration/TP_08_Api"
)

func main(){

	fmt.Println("Trabajos Prácticos LAB3 - Federico Bettic")

	fmt.Println("\n********************************************************" )
	fmt.Println("***           Ejercicio N° 3 - Hello World!          ***" )
	fmt.Println("********************************************************\n" )
	TP_03.TP_03()

	fmt.Println("\n********************************************************" )
	fmt.Println("***           Ejercicio N° 4 - Calculadora           ***" )
	fmt.Println("********************************************************\n")

	TP_04.TP_04(5,10)

	fmt.Println("\n\n********************************************************" )
	fmt.Println("***              Ejercicio N° 5 - Ping               ***" )
	fmt.Println("********************************************************\n" )

	TP_05.TP_05()


	fmt.Println("\n\n********************************************************" )
	fmt.Println("***     Ejercicio N° 6 - Calculadora Modificada      ***" )
	fmt.Println("********************************************************\n" )

	TP_06.TP_06(5,10)

	//Path Suma: 			/calc/sum?a=5&b=6
	//Path Resta: 			/calc/res?a=5&b=6
	//Path Multiplicacion:	/calc/mul?a=5&b=6
	//Path Division:		/calc/div?a=5&b=6

	fmt.Println("\n********************************************************" )
	fmt.Println("***         Ejercicio N° 8 - Api Calculadora         ***" )
	fmt.Println("********************************************************\n" )

	TP_08.TP_08()

}