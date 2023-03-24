package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jeftavares/primeiro-crud-go/src/configuration/logger"
	"github.com/jeftavares/primeiro-crud-go/src/configuration/validation"
	"github.com/jeftavares/primeiro-crud-go/src/controller/model/request"
	"github.com/jeftavares/primeiro-crud-go/src/model"
	"github.com/jeftavares/primeiro-crud-go/src/view"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func (uc *userControllerInterface) CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser Controller",
		zap.String("journey", "createUser"),
		// zap.Field{
		// 	Key:    "journey",
		// 	String: "createUser",
		// },
	)
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validade user info", err,
			zap.String("journey", "createUser"),
		)
		//fmt.Sprintf("There are some incorrect fields, error=%s", err.Error())
		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	//Utiliza a interface UserDomain, atravez do contrutor
	//Exemplo vc não consegue acessar oq é especifico do domain apenas o que esta liberado quando se cria a instancia
	//Tenta dar um model. e veja
	//ou domain.
	domain := model.NewUserDomain(userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
	)

	domainResult, err := uc.service.CreateUserServices(domain)
	if err != nil {
		logger.Error("Error trying to call CreateUser service",
			err,
			zap.String("journey", "createUser"))
		c.JSON(err.Code, err)
		return
	}

	// response := response.UserResponse{
	// 	ID:    "test",
	// 	Email: userRequest.Email,
	// 	Name:  userRequest.Name,
	// 	Age:   userRequest.Age,
	// }

	logger.Info("CreateUser controller executed successfully",
		zap.String("userId", domainResult.GetID()),
		zap.String("journey", "createUser"))

	//c.JSON(http.StatusOK, response)
	//c.String(http.StatusOK, "")
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(domainResult))
}
