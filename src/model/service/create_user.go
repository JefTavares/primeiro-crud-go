package service

import (
	"github.com/jeftavares/primeiro-crud-go/src/configuration/logger"
	"github.com/jeftavares/primeiro-crud-go/src/configuration/rest_err"
	"github.com/jeftavares/primeiro-crud-go/src/model"
	"go.uber.org/zap"
)

/*
Implementa a interface UserDomain

	Model de create user, utilizado nos controllers
*/
func (ud *userDomainService) CreateUserServices(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init CreateUser Services", zap.String("journey", "CreateUser"))

	userDomain.EncryptPassword()

	//fmt.Println(ud) ud é uma interface de *useDomainService,
	//fmt.Println(userDomain.GetPassword())
	userDomainRepository, err := ud.userRepository.CreateUser(userDomain)
	if err != nil {
		logger.Error("Error trying to call repository", err,
			zap.String("journey", "CreateUser"),
		)
		return nil, err
	}

	logger.Info(
		"CreateUser service executed sucessfully",
		zap.String("userId", userDomainRepository.GetID()),
		zap.String("journey", "CreateUser"),
	)
	return userDomainRepository, nil
}
