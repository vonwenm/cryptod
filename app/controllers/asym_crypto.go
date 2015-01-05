package controllers

import (
	"github.com/revel/revel"
)

type AsymCrypto struct {
	BaseController
}

func (a AsymCrypto) DoCrypto() revel.Result {
	return nil
}
