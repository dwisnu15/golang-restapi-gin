package token

import (
	"GinAPI/constants"
	"errors"
	"fmt"
	"time"
	"github.com/dgrijalva/jwt-go"
)

const minSecretKeySize = 32

type JWTGenerator struct {
	secretKey string
}

func NewJWTGenerator(secretKey string) (Generator, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, fmt.Errorf("%s%s",constants.InvalidTokenSize, "must be at least #{minSecretKeySize} characters")
	}
	return &JWTGenerator{secretKey}, nil
}

func (J *JWTGenerator) GenerateToken(username string, duration time.Duration) (string, error) {
	//create new payload for the registered user
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", err
	}
	//create jwt with claims (signing method and duration from payload)
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodES256, payload)
	//sign the token using your secret key
	return jwtToken.SignedString([]byte(J.secretKey))
}

func (J *JWTGenerator) VerifyToken(token string) (*Payload,error) {
	//create inner function which will be used to parse received token
	keyFunc := func(token *jwt.Token)(interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return []byte(J.secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	if err != nil {
		verify, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verify.Inner, ErrExpiredToken) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}
	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, ErrInvalidToken
	}
	return payload, nil
}

