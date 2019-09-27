package consensus

import (
	amino "github.com/tendermint/go-amino"
	"github.com/orientwalt/tendermint/types"
)

var cdc = amino.NewCodec()

func init() {
	RegisterConsensusMessages(cdc)
	RegisterWALMessages(cdc)
	types.RegisterBlockAmino(cdc)
}
