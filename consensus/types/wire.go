package types

import (
	amino "github.com/tendermint/go-amino"
	"github.com/orientwalt/tendermint/types"
)

var cdc = amino.NewCodec()

func init() {
	types.RegisterBlockAmino(cdc)
}
