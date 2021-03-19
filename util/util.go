package util

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func PanicError(err error) {
	if err != nil {
		panic(err.Error())
	}
}
//
//func DBConn() *sql.DB {
//	db, err := sql.Open("postgres", os.Getenv())
//}

func GetInt64IdFromContext(c *gin.Context) int64 {
	idParam := c.Param("id")
	//because Items id is uint, use base 10 with 32 bit
	id, _ := strconv.ParseInt(idParam, 10, 64)

	return id
}



