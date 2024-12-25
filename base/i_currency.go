package base

type ICurrency interface {
	IsValidAddress(address string, currency *Currency, opts interface{}) (bool, error)
}
