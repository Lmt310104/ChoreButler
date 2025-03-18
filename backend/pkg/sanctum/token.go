package sanctum

import (
	"errors"
	"strings"
)

type NewToken struct {
	PlainText  string `json:"plain_text"`
	HashedText string `json:"hashed_text"`
}

func (nt *NewToken) GetPlainText(ID string) string {
	return ID + "|" + nt.PlainText
}

type TokenSanctum struct {
	Crypto ICryptoSanctum
}

type ITokenSactum interface {
	CreateToken() (*NewToken, error)
	SplitToken(token string) (string, string, error)
}

func NewTokenSanctum(crypto ICryptoSanctum) ITokenSactum {
	return &TokenSanctum{
		Crypto: crypto,
	}
}

func (t *TokenSanctum) CreateToken() (*NewToken, error) {
	plainText, err := GenerateRandomString(40)
	if err != nil {
		return nil, err
	}
	hashedText := t.Crypto.HMACSHA256(plainText)
	return &NewToken{
		PlainText:  plainText,
		HashedText: hashedText,
	}, nil
}

func (t *TokenSanctum) SplitToken(token string) (string, string, error) {
	parts := strings.Split(token, "|")
	if len(parts) != 2 {
		return "", "", errors.New("invalid token")
	}
	return parts[0], t.Crypto.HMACSHA256(parts[1]), nil
}

var _ ITokenSactum = (*TokenSanctum)(nil)
