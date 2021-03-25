package middlewares

import (
	"GinAPI/models/apperrors"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"sync"
	"time"
)
//https://dev.to/jacobsngoodwin/13-gin-handler-timeout-middleware-4bhg
type writerInterface interface {
	Write(b []byte) (int, error)
	WriteHeader(code int)
	writeHeader(code int)
	Header() http.Header
	SetTimedOut()
}

type timeoutWriter struct {
	gin.ResponseWriter //writer inside of writer to prevent overwriting response bodies
	h http.Header //and headers
	wbuf bytes.Buffer // The zero value for Buffer is an empty buffer ready to use.

	mu sync.Mutex //locking writer to prevent race condition
	timedOut bool //check if a request has timed out
	wroteHeader bool //to check whether the server needs to write a header
	code int //response code
}

// Writes the response, but first makes sure there
// hasn't already been a timeout
// In http.ResponseWriter interface
func (tw *timeoutWriter) Write(b []byte)(int, error) {
	tw.mu.Lock()
	defer tw.mu.Unlock()
	if tw.timedOut {
		return 0, nil
	}

	return tw.wbuf.Write(b)
}

// In http.ResponseWriter interface
func (tw *timeoutWriter) WriteHeader(code int) {
	checkWriteHeaderCode(code)
	tw.mu.Lock()
	defer tw.mu.Unlock()
	// We do not write the header if we've timed out or written the header
	if tw.timedOut || tw.wroteHeader {
		return
	}
	tw.writeHeader(code)
}
// set that the header has been written
func (tw *timeoutWriter) writeHeader(code int) {
	tw.wroteHeader = true
	tw.code = code
}

// SetTimeOut sets timedOut field to true
func (tw *timeoutWriter) SetTimedOut() {
	tw.timedOut = true
}

func checkWriteHeaderCode(code int) {
	if code < 100 || code > 999 {
		panic(fmt.Sprintf("invalid WriteHeader code %v", code))
	}
}

func TimeoutHandler(timeout time.Duration, errTimeout *apperrors.CustomError) gin.HandlerFunc {
	return func (c *gin.Context) {
		//set gin writer as custome writer
		tw := &timeoutWriter{
			ResponseWriter: c.Writer,
			h: make(http.Header),
		}
		c.Writer = tw

		//wrap the request context with a timeout
		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)
		defer cancel()

		//update gin request context
		c.Request = c.Request.WithContext(ctx)
		finished := make(chan struct{})
		panicChan := make(chan interface{}, 1)

		go func() {
			defer func() {
				if p := recover(); p!= nil {
					panicChan <- p
				}
			}()

			c.Next() //calls subsequent middlewares and handler
			finished <- struct{}{}
		}()

		select {
			case <- panicChan:
			//if we cannot recover from panic, send internal server error
			e := apperrors.NewInternal()
			tw.ResponseWriter.WriteHeader(e.Status())
			eResponse, _ := json.Marshal(gin.H{
				"error": e,
			})
			_, err := tw.ResponseWriter.Write(eResponse)
			if err != nil {
				logrus.Errorf("error creating response")
			}

			case <- finished:
				//if finished, set headers and write response
				tw.mu.Lock()
				defer tw.mu.Unlock()

				tw.ResponseWriter.Header().Set("Content-Type", "application/json")
				tw.ResponseWriter.WriteHeader(errTimeout.Status())
				eResponse, _ := json.Marshal(gin.H{
					"error": errTimeout,
				})
				tw.ResponseWriter.Write(eResponse)
				c.Abort()
				tw.SetTimedOut()

		}
	}
}
