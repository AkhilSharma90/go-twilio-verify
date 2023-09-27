package main

import (
	"github.com/akhilsharma90/go-twilio-verify/api"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//initialize config
	app := api.Config{Router: router}

	//routes
	app.Routes()

	router.Run(":8000")
}
