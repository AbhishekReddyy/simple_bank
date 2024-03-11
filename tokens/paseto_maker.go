package tokens

import (
	"strings"
	"time"

	"aidanwoods.dev/go-paseto"
	"github.com/google/uuid"
)

type PasetoMaker struct {
	secretKey paseto.V4SymmetricKey
	implicit  []byte
}

func NewPasetoMaker() Maker {
	return &PasetoMaker{paseto.NewV4SymmetricKey(), []byte("my implicit nonce")}
}

func (maker *PasetoMaker) CreateToken(username string, duration time.Duration) (string, error) {

	token := paseto.NewToken()

	tokenID, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	token.Set("id", tokenID.String())
	token.Set("username", username)
	token.SetIssuedAt(time.Now())
	token.SetExpiration(time.Now().Add(duration))
	return token.V4Encrypt(maker.secretKey, maker.implicit), nil
}

func (maker *PasetoMaker) VerifyToken(token string) (*Payload, error) {
	parser := paseto.NewParser()
	parser.AddRule(paseto.NotExpired())
	parsedToken, err := parser.ParseV4Local(maker.secretKey, token, maker.implicit)
	if err != nil {
		if strings.Contains(err.Error(), "expired") {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}
	// construct payload from token
	payload, err := getPayloadFromToken(parsedToken)
	if err != nil {
		return nil, ErrInvalidToken
	}
	return payload, nil
}

func getPayloadFromToken(t *paseto.Token) (*Payload, error) {
	id, err := t.GetString("id")
	if err != nil {
		return nil, ErrInvalidToken
	}
	username, err := t.GetString("username")
	if err != nil {
		return nil, ErrInvalidToken
	}
	issuedAt, err := t.GetIssuedAt()
	if err != nil {
		return nil, ErrInvalidToken
	}
	expiredAt, err := t.GetExpiration()
	if err != nil {
		return nil, ErrInvalidToken
	}

	return &Payload{
		ID:        uuid.MustParse(id),
		UserName:  username,
		IssuedAt:  issuedAt,
		ExpiredAt: expiredAt,
	}, nil
}

var _ Maker = (*PasetoMaker)(nil)
