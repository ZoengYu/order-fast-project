package api

import (
	"fmt"

	db "github.com/ZoengYu/order-fast-project/db/sqlc"
	"github.com/ZoengYu/order-fast-project/token"
	util "github.com/ZoengYu/order-fast-project/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	db_service db.DBService
	tokenMaker token.Maker
	router     *gin.Engine
	config     util.Config
}

func NewServer(config util.Config, db_service db.DBService) (*Server, error) {
	tokenMaker, err := token.NewJWTMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker :%w", err)
	}
	server := &Server{
		db_service: db_service,
		tokenMaker: tokenMaker,
		config:     config,
	}

	server.setupRouter()

	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()
	v1 := router.Group("/v1")
	// allow all origins request
	router.Use(cors.Default())

	v1.POST("/user", server.createUser)
	v1.POST("/user/login", server.loginUser)

	authRoutes := v1.Group("/").Use(authMiddleware(server.tokenMaker))
	authRoutes.POST("/store", server.createStore)
	authRoutes.GET("/store/:id", server.getStore)
	authRoutes.GET("/store", server.listStoresByName)
	authRoutes.PUT("/store", server.updateStore)
	authRoutes.DELETE("/store/:id", server.delStore)

	authRoutes.POST("/store/menu", server.createStoreMenu)
	authRoutes.GET("/store/menu", server.getStoreMenu)
	authRoutes.PUT("/store/menu", server.updateStoreMenu)
	authRoutes.DELETE("/store/menu/:id", server.deleteStoreMenu)
	authRoutes.GET("/store/menu_list", server.listStoreMenu)

	authRoutes.POST("store/menu/item", server.CreateMenuItem)
	authRoutes.DELETE("store/menu/item/:id", server.DeleteMenuItem)
	authRoutes.GET("/store/menu/list_items", server.listMenuItems)
	authRoutes.PUT("/store/menu/item", server.updateMenuItem)
	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
