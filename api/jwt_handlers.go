package api

import (
	"crypto/md5"
	"encoding/hex"
	"net/http"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/razielblood/corciel_inventory_manager/types"
)

func payloadFunc(data interface{}) jwt.MapClaims {
	if v, ok := data.(*types.User); ok {
		return jwt.MapClaims{
			identityKey: v.Username,
		}
	}
	return jwt.MapClaims{}
}

func identityHandler(c *gin.Context) interface{} {
	claims := jwt.ExtractClaims(c)
	return &types.User{
		Username: claims[identityKey].(string),
	}
}

func (s APIServer) handleAuthentication(c *gin.Context) (interface{}, error) {
	var loginVals types.LoginRequest
	if err := c.ShouldBind(&loginVals); err != nil {
		return "", jwt.ErrMissingLoginValues
	}
	loginRequest := types.CreateLoginRequest(loginVals.Username, getMD5Hash(loginVals.Password))

	user, err := s.store.LoginUser(loginRequest)

	if err != nil {
		return nil, jwt.ErrFailedAuthentication
	}

	return user, nil

}

func handleAuthorization(data interface{}, c *gin.Context) bool {
	if _, ok := data.(*types.User); ok {
		return true
	}

	return false
}

func handleUnauthorized(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
	})
}

func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func (s APIServer) handleSignUp(c *gin.Context) {
	newUserReq := new(types.CreateUserRequest)

	// Call BindJSON to bind the received JSON to
	// newProduct.
	if err := c.BindJSON(newUserReq); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	newUserReq.Password = getMD5Hash(newUserReq.Password)

	user, err := s.store.CreateUser(newUserReq)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.IndentedJSON(http.StatusCreated, user)

}
