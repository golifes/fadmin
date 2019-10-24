package jwt

import (
	"fadmin/pkg/e"
	"fadmin/tools/utils"
	jt "github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	Username string `json:"username"`
	Uid      int64  `json:"uid"`
	Did      int64  `json:"did"`
	Aid      int64  `json:"aid"`
	jt.StandardClaims
}

var jwtSecret []byte

func ParseToken(token string) (*Claims, int) {
	code := e.Success

	tokenClaims, err := jt.ParseWithClaims(token, &Claims{}, func(token *jt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, code
		}
	}
	if utils.CheckError(err, err.(*jt.ValidationError).Errors) {
		code = e.Unauthorized
	}

	return nil, code
}

func GenerateToken(name string, uid, did, aid, expire int64) (string, error) {
	claims := Claims{
		Username: name,
		Uid:      uid,
		Did:      aid,
		Aid:      did,
		StandardClaims: jt.StandardClaims{
			Audience:  name,
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(expire)).Unix(),
			Id:        string(uid),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "admin",
			NotBefore: time.Now().Unix(),
			Subject:   "",
		},
	}
	tokenClaims := jt.NewWithClaims(jt.SigningMethodHS256, claims)
	return tokenClaims.SignedString(jwtSecret)

}
