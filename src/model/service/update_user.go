package service

import (
	"github.com/jeftavares/primeiro-crud-go/src/configuration/logger"
	"github.com/jeftavares/primeiro-crud-go/src/configuration/rest_err"
	"github.com/jeftavares/primeiro-crud-go/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) UpdateUser(
	userId string,
	userDomain model.UserDomainInterface,
) *rest_err.RestErr {
	logger.Info("Init UpdateUser Services", zap.String("journey", "UpdateUser"))

	err := ud.userRepository.UpdateUser(userId, userDomain)
	if err != nil {
		logger.Error("Error trying to call repository", err,
			zap.String("journey", "UpdateUser"),
		)
		return err
	}

	logger.Info(
		"UpdateUser service executed sucessfully",
		zap.String("userId", userId),
		zap.String("journey", "UpdateUser"),
	)

	return nil
}
