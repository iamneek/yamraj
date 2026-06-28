package yamraj

import "errors"

var (
	ErrSessionNotFound = errors.New("session not found")
	ErrSessionExpired  = errors.New("session expired")
	ErrSessionRevoked  = errors.New("session revoked")

	ErrTokenInvalid       = errors.New("token invalid")
	ErrTokenExpired       = errors.New("token expired")
	ErrTokenMalformed     = errors.New("token malformed")
	ErrTokenMissingClaims = errors.New("token missing claims")

	ErrPasswordTooShort   = errors.New("password too short")
	ErrPasswordHashFailed = errors.New("password hash failed")
	ErrPasswordMismatch   = errors.New("password mismatch")
	ErrHashInvalid        = errors.New("hash invalid")

	ErrUnauthorized = errors.New("unauthorized")
	ErrStoreFailure = errors.New("store failure")
)
