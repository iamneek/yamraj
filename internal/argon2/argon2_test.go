package argon2

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestArgonHash(t *testing.T) {
	cases := []struct {
		name     string
		password string
		wantErr  bool
	}{
		{"Correct Hash Check", "Yamraj1234$", false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			hashed, err := ArgonHash(tc.password, DefaultParams())
			if !tc.wantErr {
				require.NoError(t, err)
				require.True(t, strings.HasPrefix(hashed, "$argon2id$"))
			} else {
				require.Error(t, err)
			}

		})
	}

}

func TestArgonHashSaltUniqueness(t *testing.T) {
	hash1, err1 := ArgonHash("Pas5W0r#d@42", DefaultParams())
	hash2, err2 := ArgonHash("Pas5W0r#d@42", DefaultParams())
	require.NoError(t, err1)
	require.NoError(t, err2)
	require.NotEqual(t, hash1, hash2)
}

func TestVerifyHash(t *testing.T) {
	hashedPass, err := ArgonHash("Yam@r4j02", DefaultParams())
	require.NoError(t, err)
	cases := []struct {
		name      string
		hash      string
		password  string
		wantMatch bool
	}{
		{"Verify Correct Password", hashedPass, "Yam@r4j02", true},
		{"Verify Incorrect Password", hashedPass, "NotYam@r4j20", false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			verified, err := VerifyHash(tc.password, tc.hash)
			require.NoError(t, err)
			if tc.wantMatch {
				require.True(t, verified)
			} else {
				require.False(t, verified)
			}
		})
	}
}

func TestVerifyMalformedHash(t *testing.T) {
	_, err := VerifyHash("PasswordHaa", "$argon2-invalid-hash")
	require.Error(t, err)
}

func TestNeedsRehash(t *testing.T) {
    currentParams := DefaultParams()
    hash, err := ArgonHash("password123", currentParams)
    require.NoError(t, err)

    needsRehash, err := NeedRehash(hash, currentParams)
    require.NoError(t, err)
    require.False(t, needsRehash)

    weakParams := Params{Memory: 16384, Time: 1, Parallelism: 2, SaltLength: 16, KeyLength: 32}
    weakHash, err := ArgonHash("password123", weakParams)
    require.NoError(t, err)

    needsRehash2, err := NeedRehash(weakHash, currentParams)
    require.NoError(t, err)
    require.True(t, needsRehash2)
}
