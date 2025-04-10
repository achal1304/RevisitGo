package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// Define a secret key to sign the JWT
var mySigningKey = []byte("my_secret_key")

// Struct for claims
type MyClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateJWT(username string) (string, error) {
	// Create the claims for the token
	claims := MyClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "my_app",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // Token expires in 24 hours
		},
	}

	// Create a new JWT token with the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	signedToken, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func VerifyJWT(signedToken string) (*MyClaims, error) {
	// Parse the token
	token, err := jwt.ParseWithClaims(signedToken, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Check token signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", token.Header["alg"])
		}
		return mySigningKey, nil
	})

	if err != nil {
		return nil, err
	}

	// Validate and return claims
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("invalid token")
	}
}

func main() {
	// Generate JWT
	username := "john_doe"
	token, err := GenerateJWT(username)
	if err != nil {
		fmt.Println("Error generating JWT:", err)
		return
	}
	fmt.Println("Generated JWT:", token)

	// Verify JWT
	claims, err := VerifyJWT(token)
	if err != nil {
		fmt.Println("Error verifying JWT:", err)
		return
	}
	fmt.Printf("Verified JWT for user: %s\n", claims.Username)
}
