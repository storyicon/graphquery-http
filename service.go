package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/storyicon/graphquery"
)

var (
	// Port is the listening port number.
	// After the service starts, you can access Playground through 127.0.0.1:port.
	Port string
	// Debug is used to set whether the Service prints the requested log.
	Debug bool
)

// Response defines the API response format.
type Response struct {
	// Data is the carrier for returning data.
	Data interface{} `json:"data"`
	// Error records the errors in this request.
	Error string `json:"error"`
	// TimeCost recorded the time wastage of the request.
	TimeCost int64 `json:"timecost"`
}

// Start is used to start the http server
func Start() {
	if !Debug {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()
	router.POST("/", func(c *gin.Context) {
		document := c.PostForm("document")
		expression := c.PostForm("expression")
		timeStart := time.Now().UnixNano()
		conseq := graphquery.ParseFromString(document, expression)
		err := ""
		if len(conseq.Errors) > 0 {
			err = conseq.Errors[0]
		}
		c.JSON(http.StatusOK, Response{
			Data:     conseq.Data,
			TimeCost: (time.Now().UnixNano() - timeStart),
			Error:    err,
		})
	})

	address := fmt.Sprintf(":%s", Port)
	router.Run(address)
}

func main() {
	flag.StringVar(&Port, "port", "8559", "http listen port")
	flag.BoolVar(&Debug, "debug", true, "debug mode")
	flag.Parse()
	Start()
}
