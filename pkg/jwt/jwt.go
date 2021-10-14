package jwt

import (
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

var jwtSecret = []byte("test")
var jwtIssuer = "test"


type Claims struct {
	UserID int64 `json:"user_id"`
	UserName string `json:"user_name"`
	jwt.StandardClaims
}

func SetSecret(secret string)  {
	jwtSecret = []byte(secret);
}

func SetIssuer(issuer string)  {
	jwtIssuer = issuer
}

func GenerateToken(userID int64, userName string) (token string,err error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour)
	claims := Claims{
		userID,
		userName,
		jwt.StandardClaims {
			ExpiresAt : expireTime.Unix(),
			Issuer : jwtIssuer,
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = tokenClaims.SignedString(jwtSecret)
	return
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}