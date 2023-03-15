package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jeftavares/primeiro-crud-go/src/configuration/logger"
	"github.com/jeftavares/primeiro-crud-go/src/controller/routes"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("About to start user application")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//A instrução router := gin.New() criará um novo roteador gin. Os roteadores podem ser inicializados de duas
	//Maneiras, uma usando gin.New() e a outra usando gin.Default()
	//A diferença é que gin.New() inicializa um roteador sem nenhum middleware enquanto o gin.Default()
	//Inicializa o roteador com logger e middlewares de recovery
	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup)
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}
