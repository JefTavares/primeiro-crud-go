package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jeftavares/primeiro-crud-go/src/configuration/logger"
	"github.com/jeftavares/primeiro-crud-go/src/configuration/validation"
	"github.com/jeftavares/primeiro-crud-go/src/controller/model/request"
	"github.com/jeftavares/primeiro-crud-go/src/model"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func CreateUser(c *gin.Context) {
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

	domain := model.NewUserDomain(userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
	)

	if err := domain.CreateUser(); err != nil {
		c.JSON(err.Code, err)
		return
	}

	// response := response.UserResponse{
	// 	ID:    "test",
	// 	Email: userRequest.Email,
	// 	Name:  userRequest.Name,
	// 	Age:   userRequest.Age,
	// }

	logger.Info("User created sucessfully", zap.String("journey", "createUser"))

	//c.JSON(http.StatusOK, response)
	c.String(http.StatusOK, "")
}
