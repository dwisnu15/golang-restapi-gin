package middlewares

import "github.com/gin-gonic/gin"

type authHeader struct {
	Token string `header:"Authorization"`
}

type invalidArgument struct {
	Field string `json:"field"`
	Value string `json:"value"`
	Tag   string `json:"tag"`
	Param string `json:"param"`
}

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		
	}
}

