package token

import (
	"crypto"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type Token struct {
	privateEd25519Key crypto.PrivateKey
	publicEd25519Key  crypto.PublicKey
	accessExpiration  time.Duration
	refreshExpiration time.Duration
}

type Output struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func New(cfg *Config) (*Token, error) {
	token := &Token{}
	var err error

	privatePemKey := []byte(cfg.PrivatePem)
	token.privateEd25519Key, err = jwt.ParseEdPrivateKeyFromPEM(privatePemKey)
	if err != nil {
		return nil, fmt.Errorf("unable to parse Ed25519 public key: %s", err)
	}

	publicPemKey := []byte(cfg.PublicPem)
	token.publicEd25519Key, err = jwt.ParseEdPublicKeyFromPEM(publicPemKey)
	if err != nil {
		return nil, fmt.Errorf("unable to parse Ed25519 public key: %s", err)
	}

	token.accessExpiration = cfg.AccessExpiration
	token.refreshExpiration = cfg.RefreshExpiration

	return token, nil
}

type Payload struct {
	Data []byte `json:"data"`
	jwt.RegisteredClaims
}

func (token *Token) CreateTokenString(accessData, refreshData any) (*Output, error) {
	accessToken, err := token.CreateAccessToken(accessData)
	if err != nil {
		return nil, err
	}

	refreshToken, err := token.CreateRefreshToken(refreshData)
	if err != nil {
		return nil, err
	}

	return &Output{AccessToken: accessToken, RefreshToken: refreshToken}, nil
}

func (token *Token) CreateAccessToken(data any) (string, error) {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		errStr := fmt.Sprintf("error marshal access data: %s", err)
		return "", errors.New(errStr)
	}

	expireAt := jwt.NewNumericDate(time.Now().Add(token.accessExpiration))
	registeredClaims := jwt.RegisteredClaims{ExpiresAt: expireAt}
	payload := Payload{Data: dataBytes, RegisteredClaims: registeredClaims}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodEdDSA, payload)
	return jwtToken.SignedString(token.privateEd25519Key)
}

func (token *Token) CreateRefreshToken(data any) (string, error) {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		errStr := fmt.Sprintf("error marshal refresh data: %s", err)
		return "", errors.New(errStr)
	}

	expireAt := jwt.NewNumericDate(time.Now().Add(token.refreshExpiration))
	registeredClaims := jwt.RegisteredClaims{ExpiresAt: expireAt}
	payload := &Payload{Data: dataBytes, RegisteredClaims: registeredClaims}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodEdDSA, payload)
	return jwtToken.SignedString(token.privateEd25519Key)
}

const (
	invalidToken        = "invalid token"
	errorMappingPayload = "error mapping the payload"
	errorUnmarshalData  = "error unmarshalling the data"
)

func (token *Token) ExtractTokenData(tokenString string, data any) error {
	checkSigningMethod := func(jwtToken *jwt.Token) (any, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodEd25519); !ok {
			return nil, fmt.Errorf("wrong signing method: %v", jwtToken.Header["alg"])
		}
		return token.publicEd25519Key, nil
	}

	jwtToken, err := jwt.ParseWithClaims(tokenString, &Payload{}, checkSigningMethod)
	if err != nil {
		errStr := fmt.Sprintf("error: %s, token: %s", err, tokenString)
		return errors.New(errStr)
	}

	if !jwtToken.Valid {
		errStr := fmt.Sprintf("%s, token: %v", invalidToken, jwtToken)
		return errors.New(errStr)
	}

	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		errStr := fmt.Sprintf("%s: %s, token: %v", invalidToken, errorMappingPayload, jwtToken)
		return errors.New(errStr)
	}

	if err := json.Unmarshal(payload.Data, data); err != nil {
		errStr := fmt.Sprintf("%s: %s, data: %s", invalidToken, errorUnmarshalData, payload.Data)
		return errors.New(errStr)
	}

	return nil
}
