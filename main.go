package main

import (
	"github.com/gin-gonic/gin"
	"github.com/riz-it/golang-gin/controller/productcontroller"
	"github.com/riz-it/golang-gin/model"
)

func main() {
	r := gin.Default()
	model.ConnectDatabase()

	r.GET("products", productcontroller.Index)
	r.GET("products/:id", productcontroller.Show)
	r.POST("products", productcontroller.Create)
	r.PUT("products/:id", productcontroller.Update)
	r.DELETE("products/:id", productcontroller.Delete)

	r.Run()
}
