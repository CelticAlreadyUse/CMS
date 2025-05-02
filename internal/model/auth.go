package model

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type ContextAuthKey string

type UserAuthRepository interface {
	Store(ctx context.Context, req User) (int64, error)
	FindByUsername(ctx context.Context,username string )(*User,error)
}
type UserAuthUsecase interface {
	Register(ctx context.Context, req RegisterRequest) (string, error)
	Login(ctx context.Context, req LoginRequest) (string, error)
}

const BearerAuthKey ContextAuthKey = "BearerAuth"

type User struct {
	ID           int64     `gorm:"primaryKey" json:"id"`
	Username     string    `gorm:"unique;not null" json:"username"`
	HashPassword string    `gorm:"not null" json:"-"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
type LoginRequest struct {
	Username string `json:"username" binding:"required,notblank,lowercase"`
	Password string `json:"password" binding:"required,notblank"`
}
type RegisterRequest struct {
	Username string `json:"username" binding:"required,notblank,lowercase"`
	Password string `json:"password" binding:"required,notblank"`
}
type CustomClaims struct {
	UserID   int64  `json:"user_id"`
	jwt.RegisteredClaims
}
type ConfigJWT struct {
	SigningKey string
	ExpTime    string
}
