package encoding

import (
	"encoding/base64"
	"errors"
)

func Encode(bytes []byte) (string, error) {
	if len(bytes) < 1 {
		return "", errors.New("encoding: input bytes cannot be empty")
	}
	encoded := base64.RawStdEncoding.EncodeToString(bytes)
	return encoded, nil
}

func Decode(encoded string) ([]byte, error) {
	if len(encoded) < 1 {
		return nil, errors.New("decoding: input string cannot be empty")
	}
	decoded, err := base64.RawStdEncoding.DecodeString(encoded)
	if err != nil {
		return nil, errors.New("decoding: failed to decode, perhaps wrong input?")
	}
	return decoded, nil
}
