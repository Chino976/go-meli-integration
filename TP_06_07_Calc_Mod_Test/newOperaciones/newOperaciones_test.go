package newOperaciones

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDiv(t *testing.T) {
	a := 2.0
	b := 5.0

	ch := make(chan Reg)
	go Div(a,b, false, ch)

	result := <- ch

	assert.Equal(t, 0.4, result.Result, fmt.Sprintf("Fallo en la division."))

	assert.Nil(t, result.Error, fmt.Sprintf("Error not nil."))
}

func TestSum(t *testing.T){
	a := 1.0
	b := 4.0

	ch := make(chan Reg)
	go Sum(a,b, false, ch)

	result := <- ch

	assert.Equal(t, float64(5), result.Result, fmt.Sprintf("Fallo en la suma."))
}

func TestDiv0(t *testing.T){
	a := 2.0
	b := 0.0

	ch := make(chan Reg)
	go Div(a,b, false, ch)

	result := <- ch

	assert.Equal(t, float64(-1), result.Result, fmt.Sprintf("Fallo en la division."))
	assert.NotNil(t, result.Error, fmt.Sprintf("Fallo en verificacion de division por 0"))
	assert.Equal(t, "No es posible dividir por 0", result.Error.Error(), fmt.Sprintf("Fallo el mensaje al mostrar el mensaje de verificacion de division por 0"))
}

func TestRes(t *testing.T){
	a := 1.0
	b := 4.0

	ch := make(chan Reg)
	go Res(a,b,false, ch)
	result := <- ch

	assert.Equal(t, float64(-3), result.Result, fmt.Sprintf("Fallo en la resta"))
}

func TestMul(t *testing.T){

	a := 5.0
	b := 4.0

	ch := make(chan Reg)
	go Mul(a,b,false, ch)
	result := <- ch

	assert.Equal(t, float64(20), result.Result, fmt.Sprintf("Fallo en la multiplicacion"))
}