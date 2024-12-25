package base

type Opts struct {
	// TODO: maybe have sense to move to enum
	NetworkType *string

	MinLength      *int
	MaxLength      *int
	ExpectedLength *int
	Bech32Hrp      *AddressType
	AddressTypes   *AddressType
	IAddressTypes  *AddressType
	HashFunction   *string
	Regex          *string
}
