package service

import (
	"io/ioutil"
	"log"

	"github.com/dgrijalva/jwt-go"
	e "github.com/fedo3nik/GamePortal_ArticleService/internal/util/error"
	"github.com/pkg/errors"
)

const (
	refreshKeyType = "refresh"
	accessKeyType  = "access"
)

type RefreshTokenClaims struct {
	UserID    string
	KeyType   string
	CustomKey string
}

type AccessTokenClaims struct {
	UserID  string
	KeyType string
}

func (a AccessTokenClaims) Valid() error {
	if a.KeyType != accessKeyType {
		log.Println("Invalid key type")
		return e.ErrJWT
	}

	return nil
}

func (r RefreshTokenClaims) Valid() error {
	if r.KeyType != refreshKeyType {
		log.Println("Invalid key type")
		return e.ErrJWT
	}

	return nil
}

func ValidateRefreshToken(tokenString string) (id, customKey string, er error) {
	token, err := jwt.ParseWithClaims(tokenString, &RefreshTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			log.Printf("Unexpected method in token")
			return nil, e.ErrJWT
		}

		// replace .pem from file to .pem received from gRpc
		verifyBytes, err := ioutil.ReadFile("./refresh-public.pem")
		if err != nil {
			log.Printf("Unable to read public key: %v", err)
			return nil, err
		}

		verifyKey, err := jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
		if err != nil {
			log.Printf("Unable to parse public key: %v", err)
			return nil, err
		}
		return verifyKey, nil
	})

	if err != nil {
		log.Printf("Unable to parse claims: %v", err)
		return "", "", err
	}

	claims, ok := token.Claims.(*RefreshTokenClaims)
	if !ok || !token.Valid || claims.UserID == "" || claims.KeyType != "refresh" {
		log.Printf("Could not extract claims from token")
		return "", "", errors.Wrap(e.ErrJWT, "invalid token")
	}

	return claims.UserID, claims.CustomKey, nil
}

func ValidateAccessToken(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &AccessTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			log.Printf("Unexpected method in token")
			return nil, e.ErrJWT
		}

		// replace .pem from file to .pem received from gRpc
		verifyBytes, err := ioutil.ReadFile("./access-public.pem")
		if err != nil {
			log.Printf("Unable to read public key: %v", err)
			return nil, err
		}

		verifyKey, err := jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
		if err != nil {
			log.Printf("Unable to parse public key: %v", err)
			return nil, err
		}
		return verifyKey, nil
	})

	if err != nil {
		log.Printf("Unable to parse claims: %v", err)
		return "", err
	}

	claims, ok := token.Claims.(*AccessTokenClaims)
	if !ok || !token.Valid || claims.UserID == "" || claims.KeyType != accessKeyType {
		log.Printf("Could not extract claims from token")
		return "", errors.Wrap(e.ErrJWT, "invalid token")
	}

	return claims.UserID, nil
}
