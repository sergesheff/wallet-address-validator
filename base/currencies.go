package base

type Currency struct {
	Name      CurrencyTypes
	Symbol    string
	Validator func(string, *Currency, interface{}) bool

	MinLength      *int
	MaxLength      *int
	Bech32Hrp      *AddressType
	AddressTypes   *AddressType
	IAddressTypes  *AddressType
	ExpectedLength *int
	HashFunction   *string
	Regex          *string
}
