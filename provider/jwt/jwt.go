package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte("your_secret_key")

type CurrentLoginMember struct {
	ID            int    `json:"id"`
	WalletAddress string `json:"wallet_address"`
}
type Claims struct {
	CurrentLoginMember CurrentLoginMember `json:"current_login_member"`
	jwt.StandardClaims
}

// GenerateJWT 生成token
func GenerateJWT(walletAddress string, duration time.Duration) (string, error) {
	expirationTime := time.Now().Add(duration)
	claims := &Claims{
		CurrentLoginMember: CurrentLoginMember{
			WalletAddress: walletAddress,
		},
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// VerifyJWT token验证
func VerifyJWT(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	return claims, nil
}
