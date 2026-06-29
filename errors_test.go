package yamraj

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestErrSessionNotFound(t *testing.T) {
	err := ErrSessionNotFound
	require.True(t, errors.Is(err, ErrSessionNotFound))

	require.False(t, errors.Is(err, ErrSessionExpired))
	require.False(t, errors.Is(err, ErrSessionRevoked))

	require.Equal(t, "session not found", ErrSessionNotFound.Error())
}

func TestErrSessionExpired(t *testing.T) {
	err := ErrSessionExpired
	require.True(t, errors.Is(err, ErrSessionExpired))

	require.False(t, errors.Is(err, ErrSessionNotFound))
	require.False(t, errors.Is(err, ErrSessionRevoked))

	require.Equal(t, "session expired", ErrSessionExpired.Error())
}

func TestErrSessionRevoked(t *testing.T) {
	err := ErrSessionRevoked
	require.True(t, errors.Is(err, ErrSessionRevoked))

	require.False(t, errors.Is(err, ErrSessionNotFound))
	require.False(t, errors.Is(err, ErrSessionExpired))

	require.Equal(t, "session revoked", ErrSessionRevoked.Error())
}

func TestErrTokenInvalid(t *testing.T) {
	err := ErrTokenInvalid
	require.True(t, errors.Is(err, ErrTokenInvalid))

	require.False(t, errors.Is(err, ErrTokenExpired))
	require.False(t, errors.Is(err, ErrTokenMalformed))
	require.False(t, errors.Is(err, ErrTokenMissingClaims))

	require.Equal(t, "token invalid", ErrTokenInvalid.Error())
}

func TestErrTokenExpired(t *testing.T) {
	err := ErrTokenExpired
	require.True(t, errors.Is(err, ErrTokenExpired))

	require.False(t, errors.Is(err, ErrTokenInvalid))
	require.False(t, errors.Is(err, ErrTokenMalformed))
	require.False(t, errors.Is(err, ErrTokenMissingClaims))

	require.Equal(t, "token expired", ErrTokenExpired.Error())
}

func TestErrTokenMalformed(t *testing.T) {
	err := ErrTokenMalformed
	require.True(t, errors.Is(err, ErrTokenMalformed))

	require.False(t, errors.Is(err, ErrTokenInvalid))
	require.False(t, errors.Is(err, ErrTokenExpired))
	require.False(t, errors.Is(err, ErrTokenMissingClaims))

	require.Equal(t, "token malformed", ErrTokenMalformed.Error())
}

func TestErrTokenMissingClaims(t *testing.T) {
	err := ErrTokenMissingClaims
	require.True(t, errors.Is(err, ErrTokenMissingClaims))

	require.False(t, errors.Is(err, ErrTokenInvalid))
	require.False(t, errors.Is(err, ErrTokenExpired))
	require.False(t, errors.Is(err, ErrTokenMalformed))

	require.Equal(t, "token missing claims", ErrTokenMissingClaims.Error())
}

func TestErrPasswordTooShort(t *testing.T) {
	err := ErrPasswordTooShort
	require.True(t, errors.Is(err, ErrPasswordTooShort))

	require.False(t, errors.Is(err, ErrPasswordHashFailed))
	require.False(t, errors.Is(err, ErrPasswordMismatch))
	require.False(t, errors.Is(err, ErrHashInvalid))

	require.Equal(t, "password too short", ErrPasswordTooShort.Error())
}

func TestErrPasswordHashFailed(t *testing.T) {
	err := ErrPasswordHashFailed
	require.True(t, errors.Is(err, ErrPasswordHashFailed))

	require.False(t, errors.Is(err, ErrPasswordTooShort))
	require.False(t, errors.Is(err, ErrPasswordMismatch))
	require.False(t, errors.Is(err, ErrHashInvalid))

	require.Equal(t, "password hash failed", ErrPasswordHashFailed.Error())
}

func TestErrPasswordMismatch(t *testing.T) {
	err := ErrPasswordMismatch
	require.True(t, errors.Is(err, ErrPasswordMismatch))

	require.False(t, errors.Is(err, ErrPasswordTooShort))
	require.False(t, errors.Is(err, ErrPasswordHashFailed))
	require.False(t, errors.Is(err, ErrHashInvalid))

	require.Equal(t, "password mismatch", ErrPasswordMismatch.Error())
}

func TestErrHashInvalid(t *testing.T) {
	err := ErrHashInvalid
	require.True(t, errors.Is(err, ErrHashInvalid))

	require.False(t, errors.Is(err, ErrPasswordTooShort))
	require.False(t, errors.Is(err, ErrPasswordHashFailed))
	require.False(t, errors.Is(err, ErrPasswordMismatch))

	require.Equal(t, "hash invalid", ErrHashInvalid.Error())
}

func TestErrUnauthorized(t *testing.T) {
	err := ErrUnauthorized
	require.True(t, errors.Is(err, ErrUnauthorized))

	require.False(t, errors.Is(err, ErrStoreFailure))

	require.Equal(t, "unauthorized", ErrUnauthorized.Error())
}

func TestErrStoreFailure(t *testing.T) {
	err := ErrStoreFailure
	require.True(t, errors.Is(err, ErrStoreFailure))

	require.False(t, errors.Is(err, ErrUnauthorized))

	require.Equal(t, "store failure", ErrStoreFailure.Error())
}
