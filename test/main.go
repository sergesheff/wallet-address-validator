package main

import (
	"github.com/btcsuite/btcd/btcutil/bech32"
	"github.com/sergesheff/wallet-address-validator/base/curr"
)

func main() {

	decode, i, v, err := bech32.DecodeNoLimitWithVersion("abcdef1qpzry9x8gf2tvdw0s3jn54khce6mua7lmqqqxw")
	decode, i, v, err = bech32.DecodeNoLimitWithVersion("an83characterlonghumanreadablepartthatcontainsthetheexcludedcharactersbioandnumber11sg7hg6")

	if err != nil {
		return
	}

	_ = decode
	_ = i
	_ = v
	//fmt.Println("start")
	//
	//fmt.Println(crypto.HexStr2byteArray("12345688"))
	//fmt.Println(crypto.HexStr2byteArray(("12345688")))

	validator, err := curr.Currencies["Celo"].Validator("0xE37c0D48d68da5c5b14E5c1a9f1CFE802776D9FF", nil, nil)
	if err != nil {
		return
	}

	validator, err = curr.Currencies["Celo"].Validator("rDTXLQ7ZKZVKz33zJbHjgVShjsBnqMBhMN", nil, nil)
	if err != nil {
		return
	}

	_ = validator

}
