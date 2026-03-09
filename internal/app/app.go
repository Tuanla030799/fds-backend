package app

import (
	"fmt"
	"log"

	"fds-backend/internal/config"

	"github.com/gin-gonic/gin"
)

type Application struct {
	Config *config.Config
	Router *gin.Engine
}

func Run(application *Application) error {
	log.Printf("fds-backend listening on :%s", application.Config.App.Port)
	return application.Router.Run(fmt.Sprintf(":%s", application.Config.App.Port))
}
