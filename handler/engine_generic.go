package handler

import (
	"encoding/json"
	"fmt"

	"net/http"

	"gateway/models"
	"gateway/proto"
	pb "gateway/proto"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gorm.io/gorm"
)

type Request struct {
	ClientName string `json:"clientName"`
	UserId     uint32 `json:"userId"`
}
type Response struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type handler struct {
	db *gorm.DB
}

func NewHandler(db *gorm.DB) *handler {
	return &handler{
		db: db,
	}
}

func (h handler) FetchuserFromEngine(c *gin.Context) {
	ctx := c.Request.Context()
	var req Request
	err := json.NewDecoder(c.Request.Body).Decode(&req)
	if err != nil {
		log.Error().Err(err).Msg("cannot decode")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "please provide valid username, email and password"})
		return
	}
	port := FetchPortFromDB(h.db, req.ClientName)
	if port == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch port from database"})
		return
	}

	conn, engineClient, genericClient, err := GrpcConnection(port)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer conn.Close()
	if req.ClientName == "engine" {
		res, err := engineClient.FetchUser(ctx, &pb.EngineClientID{
			Id: req.UserId,
		})
		if err != nil {
			log.Printf("cannot fetch the user from engine %v", err)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "cannot fetch the user from engine"})
			return
		}
		c.JSON(200, gin.H{
			"userDetails": res,
		})
		fmt.Println(res)
		fmt.Println(req.ClientName)

	} else if req.ClientName == "generic" {

		res, err := genericClient.FetchUser(ctx, &pb.GenericClientID{
			Id: req.UserId,
		})

		if err != nil {
			log.Printf("cannot fetch the user from generic %v", err)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "cannot fetch the user from generic"})
			return
		}
		c.JSON(200, gin.H{
			"userDetails": res,
		})
		fmt.Println(res)
		fmt.Println(req.ClientName)

	} else {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid client name"})
		return
	}
}

func FetchPortFromDB(DBConn *gorm.DB, clientName string) string {
	var port models.Port
	res := DBConn.Table("ports").Where("client_name=?", clientName).Select("port").First(&port)
	if res.Error != nil {
		return "No port found"
	}

	return port.Port
}

func GrpcConnection(port string) (*grpc.ClientConn, proto.EngineRequestClient, proto.GenericRequestClient, error) {
	conn, err := grpc.NewClient("localhost:"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil, nil, err
	}
	var engineClient proto.EngineRequestClient
	var genericClient proto.GenericRequestClient

	engineClient = proto.NewEngineRequestClient(conn)

	genericClient = proto.NewGenericRequestClient(conn)

	return conn, engineClient, genericClient, nil
}
