package middlewares

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey string

func SetSecretKey(secretkey string) {
	secretKey = secretkey
}

// AuthenticationMiddleware checks if the user has a valid JWT token
func RequireAuthorization(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		fmt.Println("failed in cookie setting")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	fmt.Println("1: ", tokenString)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		fmt.Printf("algo: %s, key: %s\n", token.Method.Alg(), secretKey)
		return []byte(secretKey), nil
	})

	fmt.Println("2: ", token)
	fmt.Println("valid? ", token.Valid)
	if err != nil {
		fmt.Println("failed to parse token:", err)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	//no cheking for valid token
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
