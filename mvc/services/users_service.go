package services

import (
	"github.com/rajaseelan/golang-microservices/mvc/domain"
	"github.com/rajaseelan/golang-microservices/mvc/utils"
)

func GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userId)
}
