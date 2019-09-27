package p2p

import (
	amino "github.com/tendermint/go-amino"
	cryptoAmino "github.com/orientwalt/tendermint/crypto/encoding/amino"
)

var cdc = amino.NewCodec()

func init() {
	cryptoAmino.RegisterAmino(cdc)
}
