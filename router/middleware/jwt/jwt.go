package middleware

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/dgrijalva/jwt-go.v3"

	"github.com/ufukomer/go-boilerplate/store"
)

var (
	// Realm name to display to the user.
	Realm = "boilerplate"
	// Key is secret key used for signing.
	Key = []byte("secret key")
	// SigningAlgorithm - possible values are HS256, HS384, HS512.
	SigningAlgorithm = "HS256"
	// Timeout is the duration that a jwt token is valid.
	Timeout = time.Hour
	// MaxRefresh allows clients to refresh their token until MaxRefresh
	// has passed. Note that clients can refresh their token in the last
	// moment of MaxRefresh. This means that the maximum validity timespan
	// for a token is MaxRefresh + Timeout.
	MaxRefresh = time.Hour
	// TimeFunc provides the current time. You can override it to use another
	// time value. This is useful for testing or if your server uses a
	// different time zone than your tokens.
	TimeFunc = time.Now
)

// Authenticator authenticates if the given email and password exist.
func Authenticator(email string, password string, c *gin.Context) (string, bool) {
	_, err := store.GetUserByLogin(c, email, password)
	if err != nil {
		return email, false
	}

	return email, true
}

// Authorizator authorizes if the given email exist in db.
func Authorizator(email string, c *gin.Context) bool {
	_, err := store.GetUserByEmail(c, email)
	if err != nil {
		return false
	}

	return true
}

// Unauthorized does the opposite job of Authorizator.
func Unauthorized(c *gin.Context, code int, message string) {
	c.Header("WWW-Authenticate", "JWT realm="+Realm)
	c.Abort()

	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
	})

	return
}

// MiddlewareFunc returns Middleware interface.
func MiddlewareFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		middlewareImpl(c)
		return
	}
}

func middlewareImpl(c *gin.Context) {
	token, err := ParseToken(c)

	if err != nil {
		Unauthorized(c, http.StatusUnauthorized, err.Error())
		return
	}

	claims := token.Claims.(jwt.MapClaims)

	id := claims["id"].(string)
	c.Set("JWT_PAYLOAD", claims)
	c.Set("userID", id)

	if !Authorizator(id, c) {
		Unauthorized(c, http.StatusForbidden, "You don't have permission to access.")
		return
	}

	c.Next()
}

// ParseToken validates, and returns a token.
func ParseToken(c *gin.Context) (*jwt.Token, error) {
	var token string
	var err error

	parts := strings.Split("header:Authorization", ":")
	switch parts[0] {
	case "header":
		token, err = jwtFromHeader(c, parts[1])
	case "query":
		token, err = jwtFromQuery(c, parts[1])
	case "cookie":
		token, err = jwtFromCookie(c, parts[1])
	}

	if err != nil {
		return nil, err
	}

	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod(SigningAlgorithm) != token.Method {
			return nil, errors.New("invalid signing algorithm")
		}

		return Key, nil
	})
}

func jwtFromHeader(c *gin.Context, key string) (string, error) {
	authHeader := c.Request.Header.Get(key)

	if authHeader == "" {
		return "", errors.New("auth header empty")
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		return "", errors.New("invalid auth header")
	}

	return parts[1], nil
}

func jwtFromQuery(c *gin.Context, key string) (string, error) {
	token := c.Query(key)

	if token == "" {
		return "", errors.New("Query token empty")
	}

	return token, nil
}

func jwtFromCookie(c *gin.Context, key string) (string, error) {
	cookie, _ := c.Cookie(key)

	if cookie == "" {
		return "", errors.New("Cookie token empty")
	}

	return cookie, nil
}
