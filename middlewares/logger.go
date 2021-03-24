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

//// Log to file
func LoggerToFile() gin.HandlerFunc {
	//yet to be implemented
	//logFilePath := viper.GetString("LOG_FILE_PATH")
	//logFileName := viper.GetString("LOG_FILE_NAME")
	////log file
	//fileName := path.Join(logFilePath, logFileName)
	////write file
	//src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	//if err != nil {
	//	fmt.Println("err", err)
	//}
	//instantiation
	logger := log.New()
	//Set log level
	logger.SetLevel(log.DebugLevel)
	//Format log
	logger.SetFormatter(&log.TextFormatter{
		TimestampFormat:"2006-01-02 15:04:05",
	})
	return func(c *gin.Context) {
		// start time
		startTime := time.Now()
		// Processing request
		c.Next()
		// End time
		endTime := time.Now()
		// execution time
		latencyTime := endTime.Sub(startTime)
		// Request mode
		reqMethod := c.Request.Method
		// Request routing
		reqUri := c.Request.RequestURI
		// Status code
		statusCode := c.Writer.Status()
		// Log format
		logger.Infof("| %3d | %13v | %s | %s |",
			statusCode,
			latencyTime,
			reqMethod,
			reqUri,
		)
	}
}

