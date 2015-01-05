package controllers

import (
	"github.com/revel/revel"
)

type SymCrypto struct {
	BaseController
}

func (s SymCrypto) DoCrypto() revel.Result {
	return nil
}
