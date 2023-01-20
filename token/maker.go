package token

import "time"

type Maker interface {
	// Create new token for specific username and duration
	CreateToken(username string, duration time.Duration) (string, error)

	// Verify if token is valid or not
	VerifyToken(token string) (*Payload, error)
}
