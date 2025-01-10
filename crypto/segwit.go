package crypto

import (
	"errors"
	"strings"

	"github.com/btcsuite/btcd/btcutil/bech32"
	"github.com/sergesheff/wallet-address-validator/base"
)

type Segwit struct{}

type SegwitAddr struct {
	HRP     string
	Version byte
	Program []byte
}
type SegwitDecode struct {
	Version int
	Program []int
}

func (Segwit) convertBits(data []byte, fromBits, toBits uint8, pad bool) []byte {
	acc := uint32(0)
	bits := uint8(0)
	ret := make([]byte, 0)
	maxv := uint32((1 << toBits) - 1)

	for _, val := range data {
		if val>>fromBits != 0 {
			return nil
		}
		acc = (acc << fromBits) | uint32(val)
		bits += fromBits
		for bits >= toBits {
			bits -= toBits
			ret = append(ret, byte((acc>>bits)&maxv))
		}
	}

	if pad {
		if bits > 0 {
			ret = append(ret, byte((acc<<(toBits-bits))&maxv))
		}
	} else if bits >= fromBits || ((acc<<(toBits-bits))&maxv) != 0 {
		return nil
	}

	return ret
}

func (Segwit) Decode(addrHrp string, addr string) (*SegwitAddr, error) {
	hrp, data, version, err := bech32.DecodeNoLimitWithVersion(addr)
	if err != nil {
		return nil, err
	}

	if version == bech32.VersionUnknown || len(data) == 0 || data[0] > 16 || hrp != addrHrp {
		return nil, errors.New("invalid address")
	}

	res, err := bech32.ConvertBits(data[1:], 5, 8, false)
	if err != nil {
		return nil, err
	}

	if res == nil || len(res) < 2 || len(res) > 40 {
		return nil, errors.New("invalid address")
	}

	if data[0] == 0 && len(res) != 20 && len(res) != 32 {
		return nil, errors.New("invalid address")
	}

	if (data[0] == 0 && version == bech32.VersionM) ||
		(data[0] != 0 && version != bech32.VersionM) {
		return nil, errors.New("invalid address")
	}

	return &SegwitAddr{Version: data[0], Program: res}, nil
}

func (s Segwit) Encode(hrp string, version byte, program []byte) (*string, error) {

	var (
		ret string
		err error
	)

	if version > 0 {
		ret, err = bech32.EncodeM(hrp, s.convertBits(program, 8, 5, true))
	} else {
		ret, err = bech32.Encode(hrp, s.convertBits(program, 8, 5, true))
	}

	if err != nil {
		return nil, err
	}

	if _, err := s.Decode(hrp, ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

/////////////////////////////////////////////////////

var DefaultNetworkType = "prod"

func (s Segwit) IsValidAddress(address string, currency *base.Currency, opts *base.Opts) (bool, error) {

	if currency.Bech32Hrp == nil {
		return false, errors.New("Bech32Hrp should be provided")
	}

	networkType := "prod"

	if opts != nil && opts.NetworkType != nil {
		networkType = *opts.NetworkType
	}

	var correctBech32Hrps []string

	if networkType == "prod" || networkType == "testnet" {
		correctBech32Hrps = currency.Bech32Hrp.GetByName(networkType)
	} else {

		p := currency.Bech32Hrp.GetByName("prod")
		if p != nil {
			correctBech32Hrps = append(correctBech32Hrps, p...)
		}

		p = currency.Bech32Hrp.GetByName("testnet")
		if p != nil {
			correctBech32Hrps = append(correctBech32Hrps, p...)
		}
	}

	for _, chrp := range correctBech32Hrps {
		ret, err := s.Decode(chrp, address)
		if err != nil && ret != nil {
			en, err := s.Encode(chrp, ret.Version, ret.Program)
			if err != nil {
				return false, err
			}

			if en == nil {
				return false, nil
			}

			return *en == strings.ToLower(address), nil
		}
	}

	return false, nil
}
