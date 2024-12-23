package validators

import (
	"errors"
	"fmt"
	"github.com/sergesheff/wallet-address-validator/crypto"
	"strings"
)

type TronValidator struct {
}

func (TronValidator) decodeBase58Address(address string) ([]byte, error) {
	address = strings.TrimSpace(address)
	if len(address) != 34 {
		return nil, errors.New("address length is incorrect")
	}

	bAddress := crypto.Base58(address)
	checkSum := bAddress[len(bAddress)-4:]

	bAddress = bAddress[:len(bAddress)-4]

	hash0, err := crypto.SHA256(crypto.ByteArray2HexStr(bAddress))
	if err != nil {
		return nil, err
	}

	sha, err := crypto.SHA256(hash0)
	if err != nil {
		return nil, err
	}

	hash1 := crypto.HexStr2byteArray(sha)

	checkSum1 := hash1[:4]

	for i := 0; i < 4; i++ {
		if checkSum1[i] != checkSum[i] {
			return nil, fmt.Errorf("checksum %d does not match", i)
		}
	}

	return bAddress, nil
}

func (tv TronValidator) IsValidAddress(address string, currencyNameOrSymbol string, opts interface{}) (bool, error) {
	bAddress, err := tv.decodeBase58Address(address)
	if err != nil {
		return false, err
	}

	if bAddress == nil {
		return false, errors.New("address was not converted")
	} else if len(bAddress) != 21 {
		return false, errors.New("address length is incorrect")
	}

	const TronAddressVersionByte = 0x41

	return bAddress[0] == TronAddressVersionByte, nil

}
