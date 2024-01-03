package entities

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id" json:id`
	Name      string             `bson:"name" json:"name" validate:"required"`
	Username  string             `bson:"username" json:"username" validate:"required"`
	Password  string             `bson:"password" json:"password" validate:"required"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}

type UserRepository interface {
	InsertOne(ctx context.Context, user *User) (*User, error)
	FindOne(ctx context.Context, id string) (*User, error)
	GetAllPagination(ctx context.Context, rowPage int64, page int64, filter interface{}, setsort interface{}) ([]User, int64, error)
	UpdateOne(ctx context.Context, user *User, id string) (*User, error)
	GetByCredential(ctx context.Context, username string, password string) (*User, error)
}

type UserUsecase interface {
	InsertOne(ctx context.Context, user *User) (*User, error)
	FindOne(ctx context.Context, id string) (*User, error)
	GetAllPagination(ctx context.Context, rowPage int64, page int64, filter interface{}, setsort interface{}) ([]User, int64, error)
	UpdateOne(ctx context.Context, user *User, id string) (*User, error)
}
