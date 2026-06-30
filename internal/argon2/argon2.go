package argon2

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
	return "", nil
}

func VerifyHash(password string, hash string) (bool, error) {
	return false, nil
}

func NeedRehash(hash string, currentParams Params) (bool, error) {
	return false, nil
}