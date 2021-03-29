package controllers

import (
	"GinAPI/constants"
	"github.com/gin-gonic/gin"
	utilpass"GinAPI/util/password"
	util "GinAPI/util"
	"net/http"
)

type PasswordController struct {
	//db *sql.db
	router *gin.Engine
}

func (pcon PasswordController) SetupRouter() (*gin.Engine, error)  {
	r := gin.Default()
	pcon.router = r
	pcon.router.GET("/getpass", pcon.getPassword)
	return pcon.router, nil
}

func NewPasswordController() (*PasswordController, error) {
	controller := &PasswordController{}
	return controller, nil
}

func (pcon *PasswordController) getPassword(c *gin.Context) {
	var newPass utilpass.Password
	encode, err := utilpass.GeneratePassword()
	if err != nil {
		util.HandleFailure(c, http.StatusInternalServerError, constants.FAILED)
		return
	}
	newPass.EncodedPassword = encode
	util.HandleSuccessWithData(c, newPass)
}
