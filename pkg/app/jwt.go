/*
 * File: /pkg/app/jwt.go                                                       *
 * Project: blog-service                                                       *
 * Created At: Tuesday, 2022/06/7 , 09:55:08                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Tuesday, 2022/06/7 , 11:56:46                                *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package app

import (
	"go_start/blog_service/global"
	"go_start/blog_service/pkg/util"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	AppKey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`
	jwt.RegisteredClaims
}

func GetJWTSecret() []byte {
	return []byte(global.JWTSetting.AppSecret)
}

func GenerateToken(appKey, appSecret string) (string, error) {
	claims := Claims{
		AppKey:    util.EncodeMD5(appKey),
		AppSecret: util.EncodeMD5(appSecret),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(global.JWTSetting.Expire)),
			Issuer:    global.JWTSetting.Issuer,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(GetJWTSecret())

	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(
		token,
		&Claims{},
		func(t *jwt.Token) (interface{}, error) {
			return GetJWTSecret(), nil
		},
	)

	if tokenClaims != nil {
		claims, ok := tokenClaims.Claims.(*Claims)
		if ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
