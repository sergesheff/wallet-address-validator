package validators

import (
	"encoding/hex"
	"github.com/sergesheff/wallet-address-validator/base"
	"github.com/sergesheff/wallet-address-validator/crypto"
	"regexp"
	"slices"
	"strings"
)

type BitcoinValidator struct{}

func (BitcoinValidator) getChecksum(hashFunction string, payload string) (string, error) {
	// Each currency may implement different hashing algorithm
	switch strings.ToLower(hashFunction) {
	// blake then keccak hash chain
	case "blake256keccak256":
		b256, err := crypto.Blake2B256(payload)
		if err != nil {
			return "", err
		}

		return crypto.Keccak256Checksum(hex.EncodeToString([]byte(b256)))

	case "blake256":
		return crypto.Blacke256CheckSum(payload)

	case "keccak256":
		return crypto.Keccak256Checksum(payload)

	}

	return crypto.SHA256Checksum(payload)
}

func (bv BitcoinValidator) getAddressType(address string, currency *base.Currency) (string, error) {
	expectedLength := 25
	hashFunction := "sha256"
	regex := ""
	if currency != nil {
		if currency.ExpectedLength != nil {
			expectedLength = *currency.ExpectedLength
		}

		if currency.HashFunction != nil {
			hashFunction = *currency.HashFunction
		}

		if currency.Regex != nil {
			regex = *currency.Regex
		}
	}

	decoded := crypto.Base58(address)

	if len(decoded) > 0 {
		length := len(decoded)

		if length != expectedLength {
			return "", nil
		}

		if len(regex) > 0 {
			if !regexp.MustCompile(regex).MatchString(address) {
				return "", nil
			}
		}

		checksum := crypto.ToHex(decoded[len(decoded)-4:])
		body := crypto.ToHex(decoded[0 : len(decoded)-4])
		goodChecksum, err := bv.getChecksum(hashFunction, body)
		if err != nil {
			return "", err
		}

		if goodChecksum == checksum {
			return crypto.ToHex(decoded[0 : expectedLength-24]), nil
		}
	}

	return "", nil
}

func (bv BitcoinValidator) isValidP2PKHandP2SHAddress(address string, currency *base.Currency, opts *base.Opts) (bool, error) {
	networkType := "prod"

	if opts != nil && opts.NetworkType != nil {
		networkType = *opts.NetworkType
	}

	var correctAddressTypes []string

	if currency.AddressTypes != nil {

		addressType, err := bv.getAddressType(address, currency)
		if err != nil {
			return false, err
		}

		if len(addressType) > 0 {
			if networkType == "prod" || networkType == "testnet" {
				correctAddressTypes = currency.AddressTypes.GetByName(networkType)
			} else {
				if currency.AddressTypes.Prod != nil {
					correctAddressTypes = append(correctAddressTypes, currency.AddressTypes.Prod...)
				}

				if currency.AddressTypes.TestNet != nil {
					correctAddressTypes = append(correctAddressTypes, currency.AddressTypes.TestNet...)
				}
			}

			return slices.Contains(correctAddressTypes, addressType), nil
		}
	}

	return false, nil
}

func (bv BitcoinValidator) IsValidAddress(address string, currency *base.Currency, opts *base.Opts) (bool, error) {
	//TODO implement me
	panic("implement me")
}
