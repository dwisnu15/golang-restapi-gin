package middlewares

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"time"
)

func makeLogEntry(c *gin.Context) *log.Entry {
	if c == nil {
		return log.WithFields(log.Fields{
			"at": time.Now().Format("2006-01-02 15:04:05"),
		})
	}
	//goroutines
	//copy context to create a read-only
	ctxCpy := c.Copy()

	return log.WithFields(log.Fields{
		"time":     time.Now().Format("2006-01-02 15:04:05"),
		"method": ctxCpy.Request.Method,
		"uri":    ctxCpy.Request.URL.String(),
	})
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		makeLogEntry(c).Info("Incoming request")
	}
}

func ErrorResponse(err error) gin.H {
	return gin.H{"error : ":err.Error()}
}

//func errorHandler(c *gin.Context, err error) bool {
//	if err != nil {
//		switch err.Error(){
//
//		}
//	}
//}
