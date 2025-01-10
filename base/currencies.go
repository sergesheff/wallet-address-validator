package base

type Currency struct {
	Name      CurrencyTypes
	Symbol    string
	Validator func(string, *Currency, *Opts) (bool, error)

	MinLength      *int
	MaxLength      *int
	ExpectedLength *int
	Bech32Hrp      *AddressType
	AddressTypes   *AddressType
	IAddressTypes  *AddressType
	HashFunction   *string
	Regex          *string
}
