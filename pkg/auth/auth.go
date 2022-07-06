package auth

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTClaims struct {
	*jwt.StandardClaims
	UID      string
	UserName string
}

var (
	Secret                 = []byte("wondersafebox")
	TokenExpired     error = errors.New("Token is expired")            // token错误类型提炼
	TokenNotValidYet error = errors.New("Token not active yet")        // token错误类型提炼
	TokenMalformed   error = errors.New("That's not even a token")     // token错误类型提炼
	TokenInvalid     error = errors.New("Couldn't handle this token:") // token错误类型提炼
)

func CreateJwtToken(id string, username string) (string, error) {
	expireTime := time.Now().Add(time.Hour * 24 * 7).Unix()
	claims := JWTClaims{
		&jwt.StandardClaims{
			ExpiresAt: expireTime,
			NotBefore: time.Now().Unix() - 1000,
			Issuer:    "wonders",
		},
		id,
		username,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(Secret)
	if err != nil {
		return "", err
	}

	return signedToken, nil

}

func ValidateToken(signedToken string) (*JWTClaims, error) {

	tokenClaims, err := jwt.ParseWithClaims(signedToken, &JWTClaims{}, func(t *jwt.Token) ([]byte, error) {
		return Secret, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := tokenClaims.Claims.(*JWTClaims); ok && tokenClaims.Valid {
		return claims, nil
	}
	return nil, err
}

func SHA256Secret(password string) string {
	hashPassword := sha256.Sum256(append([]byte(password), Secret...))
	hexadecimal := hex.EncodeToString(hashPassword[:])

	return hexadecimal
}
