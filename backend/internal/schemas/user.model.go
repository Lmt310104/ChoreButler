package schemas

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/v2/bson"
	"time"
)

type User struct {
	ID          bson.ObjectID `json:"id" bson:"_id"`
	Email       string        `json:"email" bson:"email"`
	PhoneNumber string        `json:"phone_number" bson:"phone_number"`
	Password    string        `json:"password" bson:"password"`
	FullName    string        `json:"full_name" bson:"full_name"`
	Role        string        `json:"role" bson:"role"`
	CreatedAt   string        `json:"created_at" bson:"created_at"`
	UpdatedAt   string        `json:"updated_at" bson:"updated_at"`
	Children    []Children    `json:"children" bson:"children"`
}

type Admin struct {
	ID       bson.ObjectID `json:"id" bson:"_id"`
	Username string        `json:"username" bson:"username"`
	Password string        `json:"password" bson:"password"`
	FullName string        `json:"full_Name" bson:"full_name"`
	Role     string        `json:"role" bson:"role"`
}

type Children struct {
	ID        bson.ObjectID `json:"id" bson:"_id"`
	ParentID  string        `json:"parent_id" bson:"parent_id"`
	FullName  string        `json:"full_name" bson:"full_name"`
	Age       int           `json:"age" bson:"age"`
	CreatedAt string        `json:"created_at" bson:"created_at"`
	UpdatedAt string        `json:"updated_at" bson:"updated_at"`
}

type SanctumToken struct {
	ID        bson.ObjectID `bson:"_id,omitempty"`
	CreatedAt time.Time     `bson:"created_at,omitempty"`
	UpdatedAt time.Time     `bson:"updated_at,omitempty"`
	AdminID   uuid.UUID     `bson:"admin_id"`
	Token     string        `bson:"token"`
	ExpiredAt time.Time     `bson:"expired_at"`
	IsRevoked bool          `bson:"is_revoked"`
	Ability   string        `bson:"ability"`
}
