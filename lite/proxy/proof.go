package proxy

import (
	"github.com/orientwalt/tendermint/crypto/merkle"
)

func defaultProofRuntime() *merkle.ProofRuntime {
	prt := merkle.NewProofRuntime()
	prt.RegisterOpDecoder(
		merkle.ProofOpSimpleValue,
		merkle.SimpleValueOpDecoder,
	)
	return prt
}
