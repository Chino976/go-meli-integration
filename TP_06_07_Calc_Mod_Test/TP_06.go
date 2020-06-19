package TP_06

import (
	"fmt"
	"github.com/chino976/go-meli-integration/TP_06_07_Calc_Mod_Test/newOperaciones"
)

func TP_06(a, b float64){
	ch := make(chan newOperaciones.Reg)


	/*
	fmt.Printf("Ingrese Primer número: ")
	fmt.Scanln(&a)
	fmt.Printf("Ingrese Segundo número: ")
	fmt.Scanln(&b)
	*/

	newOperaciones.Operar(a, b, ch, false)
	for i := 0; i < 4; i++ {
		fmt.Printf("%+v \n", <- ch)
	}

}




