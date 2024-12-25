package validators

import (
	"fmt"
	"github.com/eknkc/basex"
	"github.com/sergesheff/wallet-address-validator/base"
	"github.com/sergesheff/wallet-address-validator/crypto"
	"regexp"
)

type RippleValidator struct{}

func (RippleValidator) getAllowedChars() string {
	return "rpshnaf39wBUDNEGHJKLM4PQRST7VWXYZ2bcdeCg65jkm8oFqi1tuvAxyz"
}

func (rv RippleValidator) IsValidAddress(address string, currency *base.Currency, opts interface{}) (bool, error) {

	rx := regexp.MustCompile(fmt.Sprintf(`^r[%s]{27,35}$`, rv.getAllowedChars()))

	if rx.MatchString(address) {
		return rv.VerifyCheckSum(address)
	}

	return false, nil
}

func (rv RippleValidator) VerifyCheckSum(address string) (bool, error) {
	codec, err := basex.NewEncoding(rv.getAllowedChars())
	if err != nil {
		return false, err
	}

	decode, err := codec.Decode(address)
	if err != nil {
		return false, err
	}

	computedChecksum, err := crypto.SHA256Checksum(crypto.ToHex(decode[0 : len(decode)-4]))
	if err != nil {
		return false, err
	}

	checksum := crypto.ToHex(decode[len(decode)-4:])

	return checksum == computedChecksum, nil

}
