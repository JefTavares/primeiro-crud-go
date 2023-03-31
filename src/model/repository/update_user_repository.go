package repository

import (
	"context"
	"os"

	"github.com/jeftavares/primeiro-crud-go/src/configuration/logger"
	"github.com/jeftavares/primeiro-crud-go/src/configuration/rest_err"
	"github.com/jeftavares/primeiro-crud-go/src/model"
	"github.com/jeftavares/primeiro-crud-go/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (ur *userRepository) UpdateUser(
	userId string,
	userDomain model.UserDomainInterface,
) *rest_err.RestErr {

	logger.Info("Init UpdateUser repository", zap.String("journey", "UpdateUser"))
	collection_name := os.Getenv(MONGODB_USER_COLLECTION)

	collection := ur.databaseConnection.Collection(collection_name)

	value := converter.ConvertDomainToEntity(userDomain)
	//Converte o userId para primitive (formato correto do mongoDB). Não precis validar o erro, já é validado no controller
	userIdHex, _ := primitive.ObjectIDFromHex(userId)

	filter := bson.D{{Key: "_id", Value: userIdHex}}
	update := bson.D{{Key: "$set", Value: value}}

	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		logger.Error("Error trying to UpdateUser repository", err, zap.String("journey", "UpdateUser"))
		return rest_err.NewInternalServerError(err.Error())
	}

	logger.Info("UpdateUser repository excuted successfully",
		zap.String("userId", userId),
		zap.String("journey", "UpdateUser"),
	)
	return nil

}
