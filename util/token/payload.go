package token

import (
	"errors"
	"github.com/google/uuid" //for uuid
	"time"
)

//i should refactor this
var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

//Payload data of the token
type Payload struct {
	ID uuid.UUID `json:"id"`
	Username string `json:"username"`
	CreatedAt time.Time `json:"created_at"`
	ExpiredAt time.Time `json:"expired_at"`
}



//Create a new token payload for the requesting username
func NewPayload(username string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	payload := &Payload{
		ID: tokenID,
		Username: username,
		CreatedAt: time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return payload, nil
}

func (p *Payload) Valid() error {
	if time.Now().After(p.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}
