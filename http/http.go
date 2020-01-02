package http

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/snlansky/glibs/logging"
)

var logger = logging.MustGetLogger("http")

type Server struct {
	router *gin.Engine
}

func New() *Server {
	router := gin.New()
	router.Use(LoggerWriter(), RecoveryWriter())

	api := router.Group("/api")
	{
		api.GET("/ping", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "pong")
		})

		api.POST("/meeting/registerFaceID", RegisterFaceID)

		//contract := api.Group("/contract")
		//{
		//
		//}
	}

	return &Server{router: router}
}

func (s *Server) Start(addr string) {
	err := s.router.Run(addr)
	if err != nil {
		panic(fmt.Sprintf("start Http server [%s] error:%v", addr, err))
	}
}
