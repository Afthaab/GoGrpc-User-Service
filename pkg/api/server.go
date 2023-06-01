package api

import (
	"log"
	"net"

	"github.com/gin-gonic/gin"
	handler "github.com/profile/service/pkg/api/handler"
	"github.com/profile/service/pkg/pb"
	"google.golang.org/grpc"
)

type ServerHttp struct {
	Engine *gin.Engine
}

func NewGRPCServer(userHandler *handler.UserHandler, grpcPort string) {
	lis, err := net.Listen("tcp", ":"+grpcPort)
	if err != nil {
		log.Fatalln("Failed to listen to GRPC Port", err)
	}

	//creating a new GRPC Server
	grpcServer := grpc.NewServer()

	pb.RegisterProfileManagementServer(grpcServer, userHandler)

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalln("Could not servre the grpc Server", err)
	}

}

func NewServerHttp(userHandler *handler.UserHandler) *ServerHttp {
	engine := gin.New()

	engine.Use(gin.Logger())

	go NewGRPCServer(userHandler, "8890")

	return &ServerHttp{
		Engine: engine,
	}
}

func (s *ServerHttp) Start() {
	s.Engine.Run(":7778")
}
