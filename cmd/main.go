package main

import (
	"github.com/axrav/antispam/pkg"
	"github.com/axrav/antispam/training"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode) // set to release modeß
	r := gin.New()
	go training.TrainModel(training.ReadDataset()) // train the model in the background so that the server can start up
	pkg.SetupRoutes(r)
	r.Run(":8080") // listen and serve on

}
