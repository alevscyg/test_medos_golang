package storage

import "medos/interlan/app/model"

// RefreshTokenRepository ...
type RefreshTokenRepository interface {
	CreateTokens(int64, string, string) (*model.Auth, error)
	RefreshAccess(int64, string, string) (*model.Auth, error)
}