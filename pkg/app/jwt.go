package app

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/go-programming-tour-book/blog-service/global"
	"github.com/go-programming-tour-book/blog-service/pkg/util"
	"time"
)

type Claims struct {
	AppKey string 	`json:"app_key"`
	AppSecret string `json:"app_secret"`
	jwt.StandardClaims
}

// 获取JWT secret
func GetJWTSecret() []byte {
	return []byte(global.JWTSetting.Secret)

}

// 根据AppKey AppSecret 以及在项目配置中设置的签发者 （Isuuer） 和过期时间（ExpiresAt）,根据指定的算法生成签名后的Token.
func GenerateToken(appKey, appSecret string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(global.JWTSetting.Expire)
	claims :=Claims{
		AppKey:         util.EncodeMD5(appKey),
		AppSecret:      util.EncodeMD5(appSecret),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer: global.JWTSetting.Issuer,
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(GetJWTSecret())
	return token,err

}

// 主要功能是解析和校验token，其流程是解析传入的token,然后根据Cliams 的相关属性要求进行校验。
func ParseToken(token string) (*Claims,error) {
	// ParseWithClaims 用于解析鉴权的声明，方法内部是具体的解码和校验的过程，返回*token
	tokenCliams, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return GetJWTSecret(), nil

	})
	if tokenCliams!=nil {
		claims,ok := tokenCliams.Claims.(*Claims)
		//Valid 验证基于时间的声明，如过期时间、签发者
		if ok && tokenCliams.Valid {
			return claims,nil

		}

	}
	return nil, err
}