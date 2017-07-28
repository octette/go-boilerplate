package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	jwt "gopkg.in/dgrijalva/jwt-go.v3"
)

// Login structure.
type login struct {
	Email    string
	Password string
}

// LoginHandler can be used by clients to get a jwt token.
// Payload needs to be json in the form of {"email": "EMAIL", "password": "PASSWORD"}.
// Reply will be of the form {"token": "TOKEN"}.
func Login(c *gin.Context) {

	var l login

	if c.Bind(&l) != nil {
		mw.unauthorized(c, http.StatusBadRequest, "Missing Email or Password")
		return
	}

	if mw.Authenticator == nil {
		mw.unauthorized(c, http.StatusInternalServerError, "Missing define authenticator func")
		return
	}

	userID, ok := mw.Authenticator(login.Email, login.Password, c)

	if !ok {
		mw.unauthorized(c, http.StatusUnauthorized, "Incorrect Email / Password")
		return
	}

	// Create the token
	token := jwt.New(jwt.GetSigningMethod(mw.SigningAlgorithm))
	claims := token.Claims.(jwt.MapClaims)

	if mw.PayloadFunc != nil {
		for key, value := range mw.PayloadFunc(login.Email) {
			claims[key] = value
		}
	}

	if userID == "" {
		userID = login.Email
	}

	expire := mw.TimeFunc().Add(mw.Timeout)
	claims["id"] = userID
	claims["exp"] = expire.Unix()
	claims["orig_iat"] = mw.TimeFunc().Unix()

	tokenString, err := token.SignedString(mw.Key)

	if err != nil {
		mw.unauthorized(c, http.StatusUnauthorized, "Create JWT Token faild")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token":  tokenString,
		"expire": expire.Format(time.RFC3339),
	})
}
