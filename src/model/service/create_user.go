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
	logger.Info("Init create user Services", zap.String("journey", "createUser"))

	userDomain.EncryptPassword()

	//fmt.Println(ud) ud é uma interface de *useDomainService,
	//fmt.Println(userDomain.GetPassword())
	userDomainRepository, err := ud.userRepository.CreateUser(userDomain)
	if err != nil {
		logger.Info("Init create model", zap.String("journey", "createUser"))
		return nil, err
	}

	return userDomainRepository, nil
}
