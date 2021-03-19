package models

import (
	//_ "github.com/spf13/viper"
	_"github.com/jinzhu/gorm"//i dont actually have to use this
)

//type Items struct {
//	ID     int64   `json:"id"` //in postgres ID is serial type
//	Name  string `json:"name"`
//	Price int `json:"price"`
//}
//
//type CreateItemInput struct {
//	Name  string `mapstructure:"name" binding:"required"`
//	Price int `mapstructure:"name" binding:"required"`
//}
//
//type UpdateItemInput struct {
//	Name  string `mapstructure:"name"`
//	Price int `mapstructure:"price"`
//}

type Items struct {
	ID     int64   `json:"id" gorm:"primary_key"` //in postgres ID is serial type
	Name  string `json:"name"`
	Price int `json:"price"`
}

type CreateItemInput struct {
	Name  string `json:"name" binding:"required" min:2 max:50`
	Price int `json:"price" binding:"required"`
}

type UpdateItemInput struct {
	Name  string `json:"name"`
	Price int `json:"price"`
}