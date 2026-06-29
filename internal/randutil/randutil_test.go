package randutil

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerateRandomBytes(t *testing.T) {
	cases := []struct {
		name    string
		size    int
		wantErr bool
	}{
		{"valid size", 10, false},
		{"zero size", 0, true},
		{"negative size", -7, true},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			bytes, err := GenerateRandomBytes(tc.size)
			if tc.wantErr {
				require.Error(t, err)
				require.Nil(t, bytes)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.size, len(bytes))
			}
		})
	}
}
