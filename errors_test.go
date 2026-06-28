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
	require.True(t, errors.Is(err,  ErrSessionExpired ))

	require.False(t, errors.Is(err, ErrSessionNotFound))
	require.False(t, errors.Is(err, ErrSessionRevoked))
	
	require.Equal(t, "session expired", ErrSessionExpired.Error())
}

func TestErrSessionRevoked(t *testing.T) {
	err := ErrSessionRevoked 
	require.True(t, errors.Is(err,  ErrSessionRevoked ))

	require.False(t, errors.Is(err, ErrSessionNotFound))
	require.False(t, errors.Is(err, ErrSessionExpired))
	
	require.Equal(t, "session revoked", ErrSessionRevoked.Error())
}


func TestErrTokenInvalid(t *testing.T) {
	err := ErrTokenInvalid
	require.True(t, errors.Is(err, ErrTokenInvalid))
	require.False(t, errors.Is(err, ErrTokenExpired))
}