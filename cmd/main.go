package main

import (
	brasilapi "api/busca-cep/services"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/:cep", brasilapi.GetAddressByCEP)

	router.Run(":8080")
}
