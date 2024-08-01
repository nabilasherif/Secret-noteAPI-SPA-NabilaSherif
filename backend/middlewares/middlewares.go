package middlewares

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/time/rate"
)

var secretKey string

var limiter = rate.NewLimiter(rate.Every(time.Minute), 20)

func SetSecretKey(secretkey string) {
	secretKey = secretkey
}

func RateLimiting(c *gin.Context) {
	if limiter.Allow() {
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusTooManyRequests)
	}
}

// AuthenticationMiddleware checks if the user has a valid JWT token
func RequireAuthorization(c *gin.Context) {
	authString := c.GetHeader("Authorization")
	fmt.Println("authString ", authString)

	forToken := strings.Split(authString, " ")

	if len(forToken) != 2 || forToken[0] != "Bearer" {
		fmt.Println("failed here in checking for token")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token, err := jwt.Parse(forToken[1], func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		fmt.Println("failed in parsing token")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username, ok := claims["username"].(string)
		if !ok {
			fmt.Println("failed to get username from claims")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Set("username", username)
	} else {
		fmt.Println("failed in key setting: token invalid or claims extraction failed")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	c.Next()
}

type MyClaim struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateToken(username, secretKey string, expirationDate time.Time) (string, error) {
	claim := MyClaim{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationDate),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
