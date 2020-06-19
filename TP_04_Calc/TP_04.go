package TP_04

import (
	"fmt"
	"github.com/chino976/go-meli-integration/TP_04_Calc/operaciones"
)

func TP_04(a,b float64) {

	var resultados operaciones.Resultados

	resultados = operaciones.Operar(a,b)

	fmt.Printf("Resultado real: " + "%+v", resultados)

	resultados =  operaciones.Desviar(resultados,2)

	fmt.Printf("\nResultado modificado: " + "%+v", resultados)

}

