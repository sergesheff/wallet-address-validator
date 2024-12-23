package crypto

import (
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"github.com/btcsuite/btcutil/base58"
	"github.com/decred/dcrd/crypto/blake256"
	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/sha3"
	"strconv"
	"strings"
)

func NumberToHex(number byte, length int) string {

	h := strconv.FormatInt(int64(number), 16)

	if len(h)%2 == 1 {
		h = "0" + h
	}

	repeatCount := length - len(h)
	if repeatCount > 0 {
		h = strings.Repeat("0", repeatCount) + h
	}

	return h
}

func byte2hexStr(b byte) string {
	hexByteMap := "0123456789ABCDEF"
	str := ""
	str += string(hexByteMap[b>>4])
	str += string(hexByteMap[b&0x0f])
	return str
}

func isHexChar(c byte) bool {
	return (c >= '0' && c <= '9') || (c >= 'A' && c <= 'F') || (c >= 'a' && c <= 'f')
}

func hexChar2byte(c byte) int {
	d := 0
	if c >= 'A' && c <= 'F' {
		d = int(c - 'A' + 10)
	} else if c >= 'a' && c <= 'f' {
		d = int(c - 'a' + 10)
	} else if c >= '0' && c <= '9' {
		d = int(c - '0')
	}
	return d
}

func ToHex(arrOfBytes []byte) string {
	var h string

	for _, value := range arrOfBytes {
		h += NumberToHex(value, -1)
	}
	return h
}

func SHA256(payload string) (string, error) {

	b, err := hex.DecodeString(payload)

	//b := make([]byte, base64.StdEncoding.DecodedLen(len(payload)))
	//	_, err := base64.StdEncoding.Decode(b, []byte(payload))

	if err != nil {
		return "", err
	}

	h := sha256.New()

	h.Write(b)
	hash := h.Sum(nil)

	//return base64.StdEncoding.EncodeToString(hash), nil
	return hex.EncodeToString(hash), nil
}

func SHA256X2(payload string) (string, error) {

	b, err := hex.DecodeString(payload)

	if err != nil {
		return "", err
	}

	h := sha256.New()

	h.Write([]byte(b))
	hash := h.Sum(nil)

	h.Reset()
	h.Write(hash)
	hash = h.Sum(nil)

	return hex.EncodeToString(hash), nil
}

func SHA256Checksum(payload string) (string, error) {
	hash, err := SHA256X2(payload)
	if err != nil {
		return "", err
	}

	return hash[0:8], nil
}

func SHA512(payload string) (string, error) {

	b, err := hex.DecodeString(payload)

	if err != nil {
		return "", err
	}

	h := sha512.New512_256()

	h.Write([]byte(b))
	hash := h.Sum(nil)

	//return base64.StdEncoding.EncodeToString(hash), nil
	return strings.ToUpper(hex.EncodeToString(hash)), nil
}

func Blacke256(payload string) (string, error) {

	b, err := hex.DecodeString(payload)

	if err != nil {
		return "", err
	}

	hash := blake256.Sum256(b)

	//return base64.StdEncoding.EncodeToString(hash), nil
	return hex.EncodeToString(hash[:]), nil
}

func Blacke256CheckSum(payload string) (string, error) {
	hash, err := Blacke256(payload)
	return hash[0:8], err
}

func Blake2B(payload string, outlen int) (string, error) {

	b, err := hex.DecodeString(payload)
	if err != nil {
		return "", err
	}

	h, err := blake2b.New(outlen, nil)
	if err != nil {
		return "", err
	}

	if _, err := h.Write(b); err != nil {
		return "", err
	}

	hash := h.Sum(nil)

	return hex.EncodeToString(hash), nil
}

func Keccak256(payload string) (string, error) {

	h := sha3.NewLegacyKeccak256()

	if _, err := h.Write([]byte(payload)); err != nil {
		return "", err
	}

	hash := h.Sum(nil)

	return hex.EncodeToString(hash), nil
}

func Keccak256Checksum(payload string) (string, error) {
	hash, err := Keccak256(payload)
	return hash[0:8], err
}

func Blake2B256(payload string) (string, error) {
	return Blake2B(payload, 32)
}

func Base58(payload string) []byte {
	return base58.Decode(payload)
}

func ByteArray2HexStr(payload []byte) string {
	var str string

	for _, b := range payload {
		str += byte2hexStr(b)
	}

	return str
}

func HexStr2byteArray(str string) []byte {
	byteArray := []byte{}
	d := 0
	j := 0

	for i := 0; i < len(str); i++ {
		c := str[i]
		if isHexChar(c) {
			d <<= 4
			d += hexChar2byte(c)
			j++
			if j%2 == 0 {
				byteArray = append(byteArray, byte(d))
				d = 0
			}
		}
	}
	return byteArray
}
