package operaciones

import (
	"errors"
	"fmt"
	"os"
)

type Resultados struct {
	suma float64
	resta float64
	division float64
	multiplicacion float64
}

func sum(a, b float64) float64{
	return a+b
}

func res(a, b float64) float64{
	return a-b
}

func div(a, b float64) (float64, error){

	if b == 0 {
		return -1, errors.New("Division por 0")
	}

	return a/b, nil
}

func mul(a, b float64) float64{
	return a*b
}

func Operar(a, b float64) Resultados{

	rDiv, err := div(a, b)

	if err != nil {
		fmt.Println("Error:",err.Error())
		os.Exit(1)
	}

	return Resultados{
		suma: sum(a,b),
		resta: res(a,b),
		division: rDiv,
		multiplicacion: mul(a,b),
	}
}

func Desviar(resultado Resultados, a float64) Resultados{

	resultado.suma += a
	resultado.resta += a
	resultado.division += a
	resultado.multiplicacion += a

	return resultado
}
