package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/razielblood/corciel_inventory_manager/storage"
)

type APIServer struct {
	listenAddr string
	listenPort string
	store      storage.Storage
}

func NewAPIServer(listenAddr, listenPort string, store storage.Storage) *APIServer {
	return &APIServer{listenAddr: listenAddr, listenPort: listenPort, store: store}
}

func (s *APIServer) Run() {
	router := gin.Default()

	router.GET("/products", s.handleGetProducts)
	router.GET("/products/:id", s.handleGetProductByID)
	router.POST("/products", s.handlePostProduct)

	router.Run(fmt.Sprintf("%v:%v", s.listenAddr, s.listenPort))
}
