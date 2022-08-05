package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	db "github.com/irossa/educate/db/sqlc"
	"github.com/irossa/educate/token"
	"github.com/irossa/educate/util"
)

type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}
	server.store = store
	//server := &Server(store: store, router: router)
	
	server.setupRouter()

	return server, nil
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.POST("/user/create", server.createUser)
	router.POST("/user/login", server.loginUser)

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	authRoutes.POST("/district/create", server.createDistrict)
	authRoutes.GET("/district", server.getDistrict)
	authRoutes.GET("/districts", server.getAllDistricts)
	authRoutes.POST("/district/update", server.updateDistrict)
	authRoutes.GET("/district/delete", server.deleteDistrict)
	
	authRoutes.GET("/user", server.getUser)
	authRoutes.GET("/users", server.getAllUsers)
	authRoutes.POST("/user/update", server.updateUser)
	authRoutes.GET("/user/delete", server.deleteUser)
	authRoutes.POST("/school/create", server.createSchool)
	authRoutes.GET("/school", server.getSchool)
	authRoutes.GET("/schools", server.getAllSchools)
	authRoutes.POST("/school/update", server.updateSchool)
	authRoutes.GET("/school/delete", server.deleteSchool)

	server.router = router
}
