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

type RefreshTokenClaims struct {
	*jwt.StandardClaims
	TokenType string
	RefreshToken
}
type RefreshToken struct {
	UserId int64 `json:"user_id"`
	Exp    int64 `json:"exp"`
}
type ExpiryToken struct {
	Token string
	Exp   int64
}

const (
	privKeyPath = "keys/jwtRS256.key"     // openssl genrsa -out app.rsa keysize
	pubKeyPath  = "keys/jwtRS256.key.pub" // openssl rsa -in app.rsa -pubout > app.rsa.pub\
)

var (
	tokenCache map[int64]map[string]ExpiryToken
)

func init() {
	tokenCache = make(map[int64]map[string]ExpiryToken)
}

func getSigningKey() (*rsa.PrivateKey, error) {
	signBytes, err := ioutil.ReadFile(privKeyPath)
	if err != nil {
		return nil, err
	}
	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		return nil, err
	}
	return signKey, nil
}

func createAccessToken(userId int64) (string, int64, error) {
	signKey, err := getSigningKey()
	if err != nil {
		return "", 0, err
	}
	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userId
	atClaims["is_admin"] = true
	exp := time.Now().Add(time.Minute * 15).Unix()
	atClaims["exp"] = exp
	at := jwt.NewWithClaims(jwt.SigningMethodRS256, atClaims)
	token, err := at.SignedString(signKey)
	if err != nil {
		return "", 0, err
	}
	return token, exp, nil
}
func createRefreshToken(userId int64) (string, int64, error) {
	signKey, err := getSigningKey()
	if err != nil {
		return "", 0, err
	}
	//Creating Refresh Token
	atRefreshClaims := jwt.MapClaims{}
	atRefreshClaims["user_id"] = userId
	exp := time.Now().Add(time.Hour * 24).Unix()
	atRefreshClaims["exp"] = exp
	atRefresh := jwt.NewWithClaims(jwt.SigningMethodRS256, atRefreshClaims)
	refreshToken, err := atRefresh.SignedString(signKey)
	if err != nil {
		return "", 0, err
	}
	return refreshToken, exp, nil
}

func createTokenPair(userId int64) (map[string]ExpiryToken, error) {
	//Creating Access Token
	token, expA, err := createAccessToken(userId)
	if err != nil {
		return nil, err
	}
	// create refresh token
	refreshToken, expR, err := createRefreshToken(userId)
	if err != nil {
		return nil, err
	}
	return map[string]ExpiryToken{
		"access_token":  ExpiryToken{token, expA},
		"refresh_token": ExpiryToken{refreshToken, expR},
	}, nil
}

func GetToken(userId int64) (map[string]ExpiryToken, error) {
	validTokenPair, containsToken := tokenCache[userId]
	now := time.Now().Unix()
	if !containsToken || (validTokenPair["access_token"].Exp <= now || validTokenPair["refresh_token"].Exp <= now) {
		var err error
		validTokenPair, err = createTokenPair(userId)
		if err != nil {
			return nil, err
		}
		tokenCache[userId] = validTokenPair
	}
	return validTokenPair, nil
}

func getPublicKey() (*rsa.PublicKey, error) {
	verifyBytes, err := ioutil.ReadFile(pubKeyPath)
	if err != nil {
		return nil, err
	}
	verifyKey, err := jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		return nil, err
	}
	return verifyKey, nil
}

func validateToken(w http.ResponseWriter, r *http.Request) (*jwt.Token, *RefreshTokenClaims, error) {
	var refreshTokenClaims RefreshTokenClaims
	token, err := request.ParseFromRequestWithClaims(r, request.OAuth2Extractor, &refreshTokenClaims, func(token *jwt.Token) (interface{}, error) {
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
	return token, &refreshTokenClaims, nil
}

func RefreshAccessToken(w http.ResponseWriter, r *http.Request) (string, error) {
	_, refreshTokenClaims, err := validateToken(w, r)
	if err != nil {
		return "", err
	}
	if refreshTokenClaims.Exp <= time.Now().Unix() {
		return "", errors.New("refresh token expired")
	}
	userId := refreshTokenClaims.UserId
	validToken, expA, err := createAccessToken(userId)
	if err != nil {
		return "", err
	}
	tokenCache[userId]["access_token"] = ExpiryToken{validToken, expA}
	return validToken, nil
}

func LogoutToken(userId int64) error {
	_, containsToken := tokenCache[userId]
	if containsToken {
		delete(tokenCache, userId)
	} else {
		return errors.New("invalid token")
	}
	return nil
}
