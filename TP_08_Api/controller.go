package TP_08

import (
	"github.com/gin-gonic/gin"
	"strconv"
)


func sum(c *gin.Context){
	a := c.Query("a")
	b := c.Query("b")
	if(a == "" || b == ""){
		c.String(400, "ERROR 400 - Variables vacias")
		return
	}
	A, err1:= strconv.ParseFloat(a, 64)
	B, err2:= strconv.ParseFloat(b, 64)
	if(err1 != nil || err2 !=nil){
		c.String(400, "ERROR 400 - Conversion de variables fallida")
		return
	}
	c.String(200, "{\"result\": %v}", A+B)
}

func res(c *gin.Context){
	a := c.Query("a")
	b := c.Query("b")

	if(a == "" || b == ""){
		c.String(400, "ERROR 400 - Variables vacias")
		return
	}
	A, err1:= strconv.ParseFloat(a, 64)
	B, err2:= strconv.ParseFloat(b, 64)
	if(err1 != nil || err2 !=nil){
		c.String(400, "ERROR 400 - Conversion de variables fallida")
		return
	}
	c.String(200, "{\"result\": %v}", A-B)
}

func mul(c *gin.Context){
	a := c.Query("a")
	b := c.Query("b")
	if(a == "" || b == ""){
		c.String(400, "ERROR 400 - Variables vacias")
		return
	}
	A, err1:= strconv.ParseFloat(a, 64)
	B, err2:= strconv.ParseFloat(b, 64)
	if(err1 != nil || err2 !=nil){
		c.String(400, "ERROR 400 - Conversion de variables fallida")
		return
	}
	c.String(200, "{\"result\": %v}", A*B)
}

func div(c *gin.Context){
	a := c.Query("a")
	b := c.Query("b")
	if(a == "" || b == ""){
		c.String(400, "ERROR 400 - Variables vacias")
		return
	}
	A, err1:= strconv.ParseFloat(a, 64)
	B, err2:= strconv.ParseFloat(b, 64)
	if(err1 != nil || err2 !=nil){
		c.String(400, "ERROR 400 - Conversion de variables fallida")
		return
	}
	if(B==0){
		c.String(400, "ERROR 400 - Se esta intentando dividir por 0")
		return
	}
	c.String(200, "{\"result\": %v}", A/B)
}


