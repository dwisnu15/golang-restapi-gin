package password

import (
	"crypto/rand"
	"encoding/base64"
	"math/big"
)

type Password struct {
	EncodedPassword string `json:"encoded_password"`
}

type PwdGenerator interface {
	GenerateRandomByte()([]byte, error)
	GeneratePassword() (string, error)
}

func GenerateRandomByte() ([]byte, error) {
	randInt, err := rand.Int(rand.Reader, big.NewInt(36))
	if err != nil {
		return nil, err
	}
	b := make([]byte, randInt.Int64())
	_, err = rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func  GeneratePassword() (string, error) {
	b, err := GenerateRandomByte()
	return base64.URLEncoding.EncodeToString(b), err
}
