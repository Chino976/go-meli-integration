package TP_08

import (
	"github.com/gin-gonic/gin"
	)

func TP_08(){
	r := gin.Default()

	r.GET("/calc/sum", sum)
	r.GET("/calc/mul", mul)
	r.GET("/calc/res", res)
	r.GET("/calc/div", div)

	r.Run()
}