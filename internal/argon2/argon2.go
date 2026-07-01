package argon2

import (
	"crypto/subtle"
	"errors"
	"fmt"
	"strings"

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

var ErrHashInvalid = errors.New("hash invalid")

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

func parseHash(encodedHash string) (Params, []byte, []byte, error) {
	splits := strings.Split(encodedHash, "$")
	if len(splits) != 6 {
		return Params{}, nil, nil, ErrHashInvalid
	}
	var hashVersion int
	if _, err := fmt.Sscanf(splits[2], "v=%d", &hashVersion); err != nil {
		return Params{}, nil, nil, ErrHashInvalid
	}
	if hashVersion != cargon2.Version {
		return Params{}, nil, nil, ErrHashInvalid
	}
	var p Params
	if _, err := fmt.Sscanf(splits[3], "m=%d,t=%d,p=%d", &p.Memory, &p.Time, &p.Parallelism); err != nil {
		return Params{}, nil, nil, ErrHashInvalid
	}
	salt, err := encoding.Decode(splits[4])
	if err != nil {
		return Params{}, nil, nil, ErrHashInvalid
	}
	hash, err := encoding.Decode(splits[5])
	if err != nil {
		return Params{}, nil, nil, ErrHashInvalid
	}
	p.KeyLength = uint32(len(hash))
	return p, salt, hash, nil
}

func VerifyHash(password string, encodedHash string) (bool, error) {
	p, salt, hash, err := parseHash(encodedHash)
	if err != nil {
		return false, err
	}
	newHash := cargon2.IDKey([]byte(password), salt, p.Time, p.Memory, p.Parallelism, p.KeyLength)
	return subtle.ConstantTimeCompare(hash, newHash) == 1, nil
}

func NeedsRehash(encodedHash string, currentParams Params) (bool, error) {
	p, _, _, err := parseHash(encodedHash)
	if err != nil {
		return false, err
	}
	return p.Memory < currentParams.Memory || p.Time < currentParams.Time ||
		p.Parallelism < currentParams.Parallelism ||
		p.KeyLength < currentParams.KeyLength, nil
}
