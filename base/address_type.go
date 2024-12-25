package base

import "strings"

type AddressType struct {
	Prod     []string
	TestNet  []string
	StageNet []string
}

func (at AddressType) GetByName(name string) []string {
	switch strings.ToLower(name) {
	case "prod":
		return at.Prod
	case "testnet":
		return at.TestNet
	case "stagenet":
		return at.StageNet
	}
	return nil
}
