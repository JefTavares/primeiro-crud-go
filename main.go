package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	mongodb "github.com/jeftavares/primeiro-crud-go/src/configuration/database/mongoDb"
	"github.com/jeftavares/primeiro-crud-go/src/configuration/logger"
	"github.com/jeftavares/primeiro-crud-go/src/controller"
	"github.com/jeftavares/primeiro-crud-go/src/controller/routes"
	"github.com/jeftavares/primeiro-crud-go/src/model/repository"
	"github.com/jeftavares/primeiro-crud-go/src/model/service"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	logger.Info("About to start user application")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	database, err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil {
		log.Fatalf(
			"Error trying to connect to database, error=%s \n",
			err.Error())
		return
	}

	userController := initDependencies(database)

	//A instrução router := gin.New() criará um novo roteador gin. Os roteadores podem ser inicializados de duas
	//Maneiras, uma usando gin.New() e a outra usando gin.Default()
	//A diferença é que gin.New() inicializa um roteador sem nenhum middleware enquanto o gin.Default()
	//Inicializa o roteador com logger e middlewares de recovery
	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, userController)
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}

func initDependencies(
	database *mongo.Database,
) controller.UserControllerInterface {
	//mongodb.InitConnection()
	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	return controller.NewUserControllerInterface(service)
}
