package argon2

import (
	"errors"
	"fmt"

	"github.com/iamneek/yamraj/internal/encoding"
	"github.com/iamneek/yamraj/internal/randutil"
	cargon2 "golang.org/x/crypto/argon2"
)

type Params struct {
	Memory      uint32
	Time        uint32
	Parallelism uint8
	SaltLength  uint32
	KeyLength   uint32
}

func DefaultParams() Params {
	return Params{
		Memory:      65536,
		Time:        1,
		Parallelism: 4,
		SaltLength:  16,
		KeyLength:   32,
	}
}

func ArgonHash(password string, params Params) (string, error) {
	salt, saltErr := randutil.GenerateRandomBytes(int(params.SaltLength))
	if saltErr != nil {
		return "", errors.New("salt generation failed")
	}
	rawHash := cargon2.IDKey([]byte(password), salt, params.Time, params.Memory, params.Parallelism, params.KeyLength)
	b64salt, enSaltErr := encoding.Encode(salt)
	if enSaltErr != nil {
		return "", errors.New("salt encoding to PHC format failed")
	}
	b64hash, enHashErr := encoding.Encode(rawHash)
	if enHashErr != nil {
		return "", errors.New("hash encoding to PHC format failed")
	}
	encodedHash := fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", cargon2.Version, params.Memory, params.Time, params.Parallelism, b64salt, b64hash)

	return encodedHash, nil
}

func VerifyHash(password string, encodedHash string) (bool, error) {
	return false, nil
}

func NeedRehash(hash string, currentParams Params) (bool, error) {
	return false, nil
}
