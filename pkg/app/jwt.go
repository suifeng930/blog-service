package app

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	AppKey string 	`json:"app_key"`
	AppSecret string `json:"app_secret"`
	jwt.StandardClaims
}
