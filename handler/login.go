package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	jwt "gopkg.in/dgrijalva/jwt-go.v3"

	"github.com/ufukomer/go-boilerplate/model"
	j "github.com/ufukomer/go-boilerplate/router/middleware/jwt"
)

// Login can be used by clients to get a jwt token.
// Payload needs to be json in the form of {"email": "EMAIL", "password": "PASSWORD"}.
// Reply will be of the form {"token": "TOKEN"}.
func Login(c *gin.Context) {

	var l model.Login

	if c.Bind(&l) != nil {
		j.Unauthorized(c, http.StatusBadRequest, "Missing Email or Password")
		return
	}

	userID, ok := j.Authenticator(l.Email, l.Password, c)

	if !ok {
		j.Unauthorized(c, http.StatusUnauthorized, "Incorrect Email / Password")
		return
	}

	// Create the token
	token := jwt.New(jwt.GetSigningMethod(j.SigningAlgorithm))
	claims := token.Claims.(jwt.MapClaims)

	if userID == "" {
		userID = l.Email
	}

	expire := j.TimeFunc().Add(j.Timeout)
	claims["id"] = userID
	claims["exp"] = expire.Unix()
	claims["orig_iat"] = j.TimeFunc().Unix()

	tokenString, err := token.SignedString(j.Key)

	if err != nil {
		j.Unauthorized(c, http.StatusUnauthorized, "Create JWT Token faild")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token":  tokenString,
		"expire": expire.Format(time.RFC3339),
	})
}

// RefreshHandler can be used to refresh a token.
// The token still needs to be valid on refresh.
// Reply will be of the form {"token": "TOKEN"}.
func Refresh(c *gin.Context) {
	token, _ := j.ParseToken(c)
	claims := token.Claims.(jwt.MapClaims)

	origIat := int64(claims["orig_iat"].(float64))

	if origIat < j.TimeFunc().Add(-j.MaxRefresh).Unix() {
		j.Unauthorized(c, http.StatusUnauthorized, "Token is expired.")
		return
	}

	// Create the token
	newToken := jwt.New(jwt.GetSigningMethod(j.SigningAlgorithm))
	newClaims := newToken.Claims.(jwt.MapClaims)

	for key := range claims {
		newClaims[key] = claims[key]
	}

	expire := j.TimeFunc().Add(j.Timeout)
	newClaims["id"] = claims["id"]
	newClaims["exp"] = expire.Unix()
	newClaims["orig_iat"] = origIat

	tokenString, err := newToken.SignedString(j.Key)

	if err != nil {
		j.Unauthorized(c, http.StatusUnauthorized, "Create JWT Token failed")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token":  tokenString,
		"expire": expire.Format(time.RFC3339),
	})
}
