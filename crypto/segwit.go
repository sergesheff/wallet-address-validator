package crypto

import (
	"errors"
	"github.com/btcsuite/btcd/btcutil/bech32"
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

func (Segwit) decode(addrHrp string, addr string) (*SegwitAddr, error) {
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
