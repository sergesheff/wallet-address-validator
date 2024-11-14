package base

type ICurrency interface {
	Validate(address string, currencyNameOrSymbol string, opts interface{}) bool
}
