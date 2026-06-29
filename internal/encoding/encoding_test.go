package encoding

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEncode(t *testing.T) {
	cases := []struct {
		name    string
		bytes   []byte
		wantErr bool
	}{
		{"Encode byte array", []byte("yamraj"), false},
		{"Encode empty byte array", []byte(""), true},
		{"Encode single byte array", []byte("a"), false},
		{"Encode bytes with all possible values [0-255]", []byte{0, 27, 127, 
			255, 128, 64}, false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			encoded, err := Encode(tc.bytes)
			if tc.wantErr {
				require.Error(t, err)
				require.Empty(t, encoded)
			} else {
				decoded, derr := Decode(encoded)
				require.NoError(t, err)
				require.NoError(t, derr)
				require.Equal(t, tc.bytes, decoded)
			}
		})
	}
}

func TestDecode(t *testing.T) {

}
