package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"log"
	"scoring_system/db"
	"scoring_system/service/routes"
)

var (
	port, mode string
)

func init() {
	flag.StringVar(&port, "port", "3000", "server listening on, default 3000")
	flag.StringVar(&mode, "mode", "debug", "server running mode, default debug mode")
}

func main() {

	db.OpenDB()
	defer db.CloseDB()

	flag.Parse()
	gin.SetMode(mode)
	router := routes.Init()

	err := router.Run(":" + port)
	if err != nil {
		log.Fatalf("Server Error: %+v", err)
	}
}
