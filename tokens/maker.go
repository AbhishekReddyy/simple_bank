package tokens

import "time"

// create interface to switch between the token makers
type Maker interface {
	// to a take a input from user and create token to return a token string
	CreateToken(username string, duration time.Duration) (string, error)

	// Create Verify to validate the given incoming created token and rerturn payload
	VerifyToken(token string) (*Payload, error)
}
