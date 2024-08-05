package main

import (
	"gateway/database"
	"gateway/handler"
	"gateway/proto"
	"log"

	"github.com/gin-gonic/gin"
)

type FetchEngineUser struct {
	proto.EngineRequestServer
}
type FetchGenericUser struct {
	proto.GenericRequestServer
}

func main() {
	db, err := database.ConnectToDatabase()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	handlerLayer := handler.NewHandler(db)

	router := gin.New()
	
	router.GET("/fetchUser/From/EngineOrGeneric", handlerLayer.FetchuserFromEngine)
	if err := router.Run(); err != nil {
		log.Fatalf("Failed to run Gin server: %v", err)
	}
}
