package main

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/leshachaplin/grpc-server-library/protocol"
	"github.com/leshachaplin/http-server-library/internal/auth"
	"github.com/leshachaplin/http-server-library/internal/config"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	cfg := config.NewConfig()

	opts := grpc.WithInsecure()
	clientConnInterface, err := grpc.Dial(cfg.GrpcPort, opts)
	if err != nil {
		log.Fatal(err)
	}
	defer clientConnInterface.Close()
	client := protocol.NewBookServiceClient(clientConnInterface)


	lib := auth.NewHandler(client)

	e := echo.New()

	e.POST("/addBook", lib.AddBook)
	e.POST("/deleteBook", lib.DeleteBook)
	e.POST("/getByAuthor", lib.GetBookByAuthor)
	e.POST("/getByName", lib.GetBookByName)
	e.GET("/getAll", lib.GeAlltBooks)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cfg.Port)))

}