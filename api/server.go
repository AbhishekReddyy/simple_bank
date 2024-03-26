package api

import (
	db "simplebank/db/sqlc"
	"simplebank/tokens"
	"simplebank/util"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker tokens.Maker
	router     *gin.Engine
}

func NewServer(config util.Config, store db.Store) *Server {
	tokenMaker := tokens.NewPasetoMaker()
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	server.setupRouter()
	return server
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.POST("/user", server.createUser)
	router.POST("/user/login", server.loginUser)

	authRoutes := router.Group("/", authMiddleware(server.tokenMaker))

	authRoutes.GET("/user/:username", server.getUser)
	authRoutes.POST("/accounts", server.createAccount)
	authRoutes.GET("/accounts/:id", server.getAccount)
	authRoutes.GET("/accounts", server.listAccount)
	authRoutes.DELETE("/accounts/:id", server.deleteAccount)
	authRoutes.POST("/transfers", server.createTransfer)

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
