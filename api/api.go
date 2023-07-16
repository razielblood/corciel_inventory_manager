package api

import (
	"fmt"
	"log"
	"os"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/razielblood/corciel_inventory_manager/storage"
)

var identityKey string = "id"

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

	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:           "test zone",
		Key:             []byte(os.Getenv("CORCIEL_INVENTORY_JWT_SECRET")),
		Timeout:         time.Hour,
		MaxRefresh:      2 * time.Hour,
		IdentityKey:     identityKey,
		PayloadFunc:     payloadFunc,
		IdentityHandler: identityHandler,
		Authenticator:   s.handleAuthentication,
		Authorizator:    handleAuthorization,
		Unauthorized:    handleUnauthorized,
		TokenLookup:     "header: Authorization, query: token, cookie: jwt",
		TokenHeadName:   "Bearer",
		TimeFunc:        time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	errInit := authMiddleware.MiddlewareInit()

	if errInit != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}

	router.POST("/login", authMiddleware.LoginHandler)
	router.POST("/sign-up", s.handleSignUp)
	router.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	auth := router.Group("/auth")

	auth.GET("/refresh_token", authMiddleware.RefreshHandler)

	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/products", s.handleGetProducts)
		auth.GET("/products/:id", s.handleGetProductByID)
		auth.PUT("/products/:id", s.handlePutProduct)
		auth.POST("/products", s.handlePostProduct)

		auth.GET("/manufacturers", s.handleGetManufacturers)
		auth.GET("/manufacturers/:id", s.handleGetManufacturerByID)
		auth.PUT("/manufacturers/:id", s.handlePutManufacturer)
		auth.POST("/manufacturers", s.handlePostManufacturer)
	}

	router.Run(fmt.Sprintf("%v:%v", s.listenAddr, s.listenPort))
}
