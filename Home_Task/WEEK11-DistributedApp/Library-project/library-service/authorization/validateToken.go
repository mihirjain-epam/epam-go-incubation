package authorization

import (
	"crypto/rsa"
	"errors"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
)

const ( // openssl genrsa -out app.rsa keysize
	pubKeyPath = "keys/jwtRS256.key.pub" // openssl rsa -in app.rsa -pubout > app.rsa.pub
)

var (
	verifyKey *rsa.PublicKey
)

func getPublicKey() (*rsa.PublicKey, error) {
	verifyBytes, err := ioutil.ReadFile(pubKeyPath)
	if err != nil {
		return nil, err
	}
	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		return nil, err
	}
	return verifyKey, nil
}

type AccessTokenClaims struct {
	*jwt.StandardClaims
	TokenType string
	AccessToken
}

type RefreshToken struct {
	UserId int64 `json:"user_id"`
}
type AccessToken struct {
	UserId     int64 `json:"user_id"`
	Authorized bool  `json:"authorized"`
	IsAdmin    bool  `json:"is_admin"`
	Exp        int64 `json:"exp"`
}

func ValidateToken(w http.ResponseWriter, r *http.Request) (*jwt.Token, *AccessTokenClaims, error) {
	var accessTokenClaims AccessTokenClaims
	token, err := request.ParseFromRequestWithClaims(r, request.OAuth2Extractor, &accessTokenClaims, func(token *jwt.Token) (interface{}, error) {
		// since we only use the one private key to sign the tokens,
		// we also only use its public counter part to verify
		publicKey, pubErr := getPublicKey()
		if pubErr != nil {
			return nil, pubErr
		}
		return publicKey, nil
	})
	// If the token is missing or invalid, return error
	if err != nil {
		return nil, nil, err
	}
	if !token.Valid {
		return nil, nil, errors.New("invalid token")
	}
	now := time.Now().Unix()
	if !accessTokenClaims.Authorized || accessTokenClaims.Exp <= now || !accessTokenClaims.IsAdmin {
		return nil, nil, errors.New("unauthorized")
	}
	return token, &accessTokenClaims, nil
}
