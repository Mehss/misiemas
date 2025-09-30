package utils

import (
	"errors"
	"fmt"
	"os"
	"time"
	model "tripatra-dct-service-config/database/model/user"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte(os.Getenv("SECRET_KEY"))

type Claims struct {
	UserID     string `json:"user_id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Role       string `json:"role"`
	Type       string `json:"type"`
	VendorID   string `json:"vendor_id"`
	VendorCode string `json:"vendor_code"`
	jwt.StandardClaims
}

func GenerateJWT(user *model.User) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID:     fmt.Sprintf("%d", user.ID),
		Email:      user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Check if Role is not nil and set the role name accordingly

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ValidateJWT(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	return claims, nil
}

func ValidateJWTMicroservice(tokenString string) (jwt.MapClaims, error) {
	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		return nil, errors.New("SECRET_KEY is not set")
	}

	// Parse token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	// Extract claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	// Get the service name from claims
	service, exists := claims["service"]
	if !exists {
		return nil, errors.New("unauthorized service: missing service claim")
	}

	// List of allowed microservices
	allowedServices := map[string]bool{
		"tap-account-payable-service": true,
		"tap-sap-integration-service": true,
		"tripatra-dct-service-config": true,
	}

	// Check if the service is in the allowed list
	if !allowedServices[service.(string)] {
		return nil, errors.New("unauthorized service: access denied")
	}

	return claims, nil
}

// GenerateJWTFromClaims generates a JWT token from claims
func GenerateJWTFromClaims(claims *Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// // Claims structure for eproc users
// type EprocClaims struct {
// 	Email string `json:"email"`
// 	Name  string `json:"name"`
// 	Role  string `json:"role"`
// 	Type  string `json:"type"`
// 	jwt.StandardClaims
// }

// // GenerateEprocJWT generates a JWT token for eproc users
// func GenerateEprocJWT(user *model.UserEproc) (string, error) {
// 	expirationTime := time.Now().Add(24 * time.Hour)

// 	claims := &Claims{
// 		UserID: user.VendorUserName, // or a unique identifier if available
// 		Name:   user.VendorDetailName,
// 		Email:  user.VendorUserName, // assuming VendorUserName is the email
// 		Role:   "vendor",
// 		Type:   "eproc",
// 		StandardClaims: jwt.StandardClaims{
// 			ExpiresAt: expirationTime.Unix(),
// 		},
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	tokenString, err := token.SignedString(jwtKey)
// 	if err != nil {
// 		return "", err
// 	}

// 	return tokenString, nil
// }
