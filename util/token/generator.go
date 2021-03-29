package token

import "time"

type Generator interface {

	GenerateToken(username string, duration time.Duration)(string, error)

	VerifyToken(token string) (*Payload, error)
}
