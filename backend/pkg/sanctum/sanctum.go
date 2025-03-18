package sanctum

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"time"
)

type TokenPayload struct {
	AdminID   uuid.UUID `json:"admin_id"`
	TokenID   int64     `json:"token_id"`
	ExpireAt  time.Time `json:"expire_at"`
	CreatedAt time.Time `json:"created_at"`
	Ability   string    `json:"ability"`
}

type Token struct {
	Token  ITokenSactum
	Crypto ICryptoSanctum
	db     *mongo.Database
}

func NewSanctumToken(token ITokenSactum, crypto ICryptoSanctum, db *mongo.Database) *Token {
	return &Token{
		Token:  token,
		Crypto: crypto,
		db:     db,
	}
}

//func (st *SanctumToken) CreateSanctumToken(adminID uuid.UUID, duration time.Duration) (string, error) {
//	// Create a new token
//	newToken, err := st.Token.CreateToken()
//	if err != nil {
//		return "", err
//	}
//
//	// Create the current timestamp
//	now := time.Now().UTC()
//
//	// First insert the token to the database to get the token ID
//	token := {
//		AdminID:   adminID,
//		Ability:   "*",
//		ExpiredAt: now.Add(duration),
//		Token:     newToken.HashedText,
//	}
//	if err := st.db.Create(&token).Error; err != nil {
//		return "", err
//	}
//
//	// Return the plain text token
//	return newToken.GetPlainText(strconv.FormatInt(token.ID, 10)), nil
// }
