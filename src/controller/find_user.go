package controller

import (
	"net/http"
	"net/mail"

	"github.com/gin-gonic/gin"
	"github.com/jeftavares/primeiro-crud-go/src/configuration/logger"
	"github.com/jeftavares/primeiro-crud-go/src/configuration/rest_err"
	"github.com/jeftavares/primeiro-crud-go/src/view"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) FindUserByID(c *gin.Context) {
	logger.Info("Init FindUserByID controller", zap.String("journey", "findUserByID"))

	//pega o user id da url, la das rotas
	userID := c.Param("userId")

	//primitive.ObjectIDFromHex, utiliza a lib do mongoDB para converter o object ID que vem na request para o object id do mongoDB
	if _, err := primitive.ObjectIDFromHex(userID); err != nil {
		logger.Error("Error trying to validade userId",
			err,
			zap.String("journey", "findUserByID"))
		errorMessage := rest_err.NewBadRequestError("UserID is not a valid id")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserByIDServices(userID)
	if err != nil {
		logger.Error("Error trying to call FindUserByID services",
			err,
			zap.String("journey", "findUserByID"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("FindUserByID controller executed successfully", zap.String("journey", "findUserByID"))
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))

}

func (uc *userControllerInterface) FindUserByEmail(c *gin.Context) {

	logger.Info("Init FindUserByEmail controller", zap.String("journey", "FindUserByEmail"))

	//pega o user id da url, la das rotas
	userEmail := c.Param("userEmail")

	if _, err := mail.ParseAddress(userEmail); err != nil {
		logger.Error("Error trying to validade userEmail",
			err,
			zap.String("journey", "FindUserByEmail"))
		errorMessage := rest_err.NewBadRequestError("userEmail is not a valid id")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserByEmailServices(userEmail)
	if err != nil {
		logger.Error("Error trying to call FindUserByEmail services",
			err,
			zap.String("journey", "FindUserByEmail"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("FindUserByEmail controller executed successfully", zap.String("journey", "FindUserByEmail"))
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))

}
