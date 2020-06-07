package main

import (
	"fmt"
	""
)

func main() {

	var resultados Resultados

	resultados = Operar(5,0)

	fmt.Printf("Resultado real: " + "%+v", resultados)

	resultados =  Desviar(resultados,2)

	fmt.Printf("\nResultado modificado: " + "%+v", resultados)

}

