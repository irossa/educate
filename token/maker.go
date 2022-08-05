package token

import "time"

//Maker is an interface for managing tokens
type Maker interface {
	//CreateToken creates a new token for a specific usaname and duration
	CreateToken(username string, duration time.Duration) (string, error)

	//VerifyToken checks if the tokin is valid or not
	VerifyToken(token string) (*Payload, error)
}
