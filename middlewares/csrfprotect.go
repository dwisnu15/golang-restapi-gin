package middlewares

import "github.com/gin-gonic/gin"

//should i use dependencies from others
//or create my own csrf validation?
const (
	CSRFToken = "X-CSRF-Token"
	CSRFKey = "csrf" //used as secret key

)

//Cross Site Request Forgery
type CSRFConfig struct {
	TokenLookup string
	ContextKey string
	//ErrorFunc func(c *gin.Context)
}

func initCSRF() {

}


func CreateCSRF() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func errorFunc(c *gin.Context) {

}