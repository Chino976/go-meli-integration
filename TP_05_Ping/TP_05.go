package TP_05

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func TP_05(){

	Resp, err := http.Get("https://api.mercadolibre.com/sites/MLA")

	fmt.Println("Error",err)

	if err != nil {
		fmt.Errorf("Error",err.Error())
		return
	}

	defer Resp.Body.Close()

	data, err := ioutil.ReadAll(Resp.Body)

	var Extract Getml

	Er := json.Unmarshal(data, &Extract)

	if Er != nil {
		fmt.Errorf("Error", Er.Error())
		return
	}

	fmt.Printf("Obtenido: " + "%+v", Extract)

}

type Getml struct {
	Id string
	Name string
	Country_id string
	Currencies[] Currencies
}

type Currencies struct {
	Id string
	Symbol string
}

