package repository

import "github.com/5MofDream/apollo/app/models"

type Example struct {
	Base
}

func (e *Example) GetList()[]models.Users{
	userList := make([]models.Users , 1)
	e.DB().Where("id <= ?" , 2).Find(&userList)
	return userList
}
