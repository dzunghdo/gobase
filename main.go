package main

import (
	"gobase/ws"
	"os"

	"gobase/config"
	"gobase/controllers/middlewares"
	"gobase/controllers/routes"
	"gobase/db"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func main() {
	config.LoadConfig()
	initLoggerAndTracer()

	_, err := db.GetMySQLConnector().Connect()
	if err != nil {
		panic(err)
	}

	db.GetRedisConnector().Connect()

	startWSServer()
	registerRoutes()
}

func registerRoutes() {
	router := gin.Default()

	router.Use(middlewares.RecoverApp())

	apiRouter := router.Group("/api")
	routes.HandleBaseRoutes(apiRouter)

	ws := ws.GetWSServer()
	apiRouter.GET("/ws", ws.HandleConnection) // no auth

	v1 := apiRouter.Group("/v1")
	v1.Use(middlewares.ResponseMiddleware)
	routes.HandleAuthRoutes(v1)

	v1Auth := v1.Group("")
	v1Auth.Use(middlewares.Authenticate)

	userRouter := v1Auth.Group("/users")
	routes.HandleUserRoutes(userRouter)

	router.Run(":" + config.GetConfig().App.Port)
}

func initLoggerAndTracer() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)

	log.WithFields(log.Fields{"string": "foo", "int": 1, "float": 1.1}).
		Info("My first event from golang to stdout")
}

func startWSServer() {
	ws := ws.GetWSServer()
	go ws.Start()
}
