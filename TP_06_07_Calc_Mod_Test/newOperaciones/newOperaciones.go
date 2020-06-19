package newOperaciones

import "errors"

type Reg struct {
	Operacion	string
	N1			float64
	N2      	float64
	Result      float64
	Desviado	bool
	Error       error
}

func Operar(a, b float64, ch chan Reg, desviar bool){

	go Sum(a, b, false, ch)
	go Res(a, b, false, ch)
	go Mul(a, b, false,ch)
	go Div(a, b, false,ch)

	if(desviar == true){
		go Sum(a+2, b, desviar,ch)
		go Res(a+2, b, desviar,ch)
		go Mul(a+2, b, desviar,ch)
		go Div(a+2, b, desviar,ch)
	}
}

func Sum(a, b float64, desviar bool, ch chan Reg){

	ch <- Reg{

		Operacion:	"Suma" + texto_desviacion(desviar),
		N1:    		a,
		N2:      	b,
		Result:   	a+b,
		Desviado: 	desviar,
		Error:	 	nil,
	}
}

func Res(a, b float64, desviar bool, ch chan Reg){

	ch <- Reg{
		Operacion:  "resta" + texto_desviacion(desviar),
		N1:      	a,
		N2:      	b,
		Result:    	a-b,
		Desviado: 	desviar,
		Error: 		nil,
	}
}

func Div(a, b float64, desviar bool, ch chan Reg){
	if b == 0 {
		ch <- Reg{
			Operacion:	"Division",
			N1:      	a,
			N2:      	b,
			Result:     -1,
			Desviado: 	desviar,
			Error:		errors.New("No es posible dividir por 0"),
		}
		return
	}

	ch <- Reg{
		Operacion:	"Division" + texto_desviacion(desviar),
		N1:      	a,
		N2:      	b,
		Result:     a/b,
		Desviado: 	desviar,
		Error: 		nil,
	}

}

func Mul(a, b float64, desviar bool, ch chan Reg) {

	ch <- Reg{
		Operacion:	"Multiplication" + texto_desviacion(desviar),
		N1:      	a,
		N2:      	b,
		Result:     a*b,
		Desviado: 	desviar,
		Error: 		nil,
	}

}

func texto_desviacion(desviar bool) string {
	if desviar == true {
		return " con desvio"
	}
	return " sin desvio"
}


