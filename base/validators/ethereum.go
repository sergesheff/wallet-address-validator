package validators

import (
	"github.com/sergesheff/wallet-address-validator/crypto"
	"regexp"
	"strconv"
	"strings"
)

type EthereumValidator struct{}

func (rv EthereumValidator) IsValidAddress(address string, currencyNameOrSymbol string, opts interface{}) (bool, error) {

	// Check if it has the basic requirements of an address
	if !regexp.MustCompile("^0x[0-9a-fA-F]{40}$").MatchString(address) {
		return false, nil
	}

	// If it's all small caps or all caps, return true
	if regexp.MustCompile("^0x[0-9a-f]{40}$").MatchString(address) ||
		regexp.MustCompile("^0x?[0-9A-F]{40}$").MatchString(address) {
		return true, nil
	}

	// Otherwise check each case
	return rv.VerifyCheckSum(address)

}

func (rv EthereumValidator) VerifyCheckSum(address string) (bool, error) {
	address = strings.TrimPrefix(address, "0x")

	addressHash, err := crypto.Keccak256(strings.ToLower(address))
	if err != nil {
		return false, err
	}

	for i := 0; i < 40; i++ {
		// The nth letter should be uppercase if the nth digit of casemap is 1
		letter := string(address[i])
		letterHash := string(addressHash[i])

		parseInt, err := strconv.ParseInt(letterHash, 16, 32)
		if err != nil {
			return false, err
		}

		if (parseInt > 7 && strings.ToUpper(letter) != letter) ||
			(parseInt <= 7 && strings.ToLower(letter) != letter) {
			return false, nil
		}
	}

	return true, nil
}
