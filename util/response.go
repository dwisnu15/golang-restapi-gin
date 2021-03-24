package util

import (
	"GinAPI/constants"
	"GinAPI/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

//handle successful request with no return data
func HandleSuccess(c *gin.Context, message string) {
	var data = models.ResponseMessage{
		Success: true,
		Message: message,
	}
	c.JSON(http.StatusOK, data)
}

//with data
func HandleSuccessWithData(c *gin.Context, data interface{}) {
	var returnData = models.ResponseBody{
		Success: true,
		Message: constants.SUCCESS,
		Data: data,
	}
	c.JSON(http.StatusOK, returnData)
}

func HandleFailure(c *gin.Context, errcode int, message string) {
	var response = models.ResponseMessage{
		Success: false,
		Message: message,
	}
	c.JSON(errcode, response)
}