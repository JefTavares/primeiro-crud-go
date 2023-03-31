package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jeftavares/primeiro-crud-go/src/configuration/logger"
	"github.com/jeftavares/primeiro-crud-go/src/configuration/rest_err"
	"github.com/jeftavares/primeiro-crud-go/src/configuration/validation"
	"github.com/jeftavares/primeiro-crud-go/src/controller/model/request"
	"github.com/jeftavares/primeiro-crud-go/src/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) UpdateUser(c *gin.Context) {
	logger.Info("Init UpdateUser Controller",
		zap.String("journey", "UpdateUser"),
	)
	var userRequest request.UserUpdateRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validade user info", err,
			zap.String("journey", "UpdateUser"),
		)
		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	userId := c.Param("userId")
	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		errRest := rest_err.NewBadRequestError("Invalid userId, must be a hex value")
		c.JSON(errRest.Code, errRest)
	}

	domain := model.NewUserUpdateDomain(
		userRequest.Name,
		userRequest.Age,
	)

	err := uc.service.UpdateUser(userId, domain)
	if err != nil {
		logger.Error("Error trying to call UpdateUser service",
			err,
			zap.String("journey", "UpdateUser"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("UpdateUser controller executed successfully",
		zap.String("userId", userId),
		zap.String("journey", "UpdateUser"))

	//c.JSON(http.StatusOK, view.ConvertDomainToResponse(domainResult))
	c.Status(http.StatusOK)
}
