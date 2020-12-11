package types

import (
	"bytes"
	"sort"

	kv "github.com/tendermint/tendermint/libs/kv"
	// "github.com/tendermint/tmlibs/common"
)

//------------------------------------------------------------------------------

// ValidatorUpdates is a list of validators that implements the Sort interface
type ValidatorUpdates []ValidatorUpdate

var _ sort.Interface = (ValidatorUpdates)(nil)

// All these methods for ValidatorUpdates:
//    Len, Less and Swap
// are for ValidatorUpdates to implement sort.Interface
// which will be used by the sort package.
// See Issue https://github.com/tendermint/abci/issues/212

func (v ValidatorUpdates) Len() int {
	return len(v)
}

// XXX: doesn't distinguish same validator with different power
func (v ValidatorUpdates) Less(i, j int) bool {
	return bytes.Compare(v[i].PubKey.Data, v[j].PubKey.Data) <= 0
}

func (v ValidatorUpdates) Swap(i, j int) {
	v1 := v[i]
	v[i] = v[j]
	v[j] = v1
}

// added by junying, 2020-12-09
func GetEventByKey(events []Event, key string) (Event, kv.Pair, bool) {
	// tags []common.KVPair
	for _, event := range events {
		for _, attribute := range event.GetAttributes() {
			if bytes.Equal(attribute.GetKey(), []byte(key)) {
				return event, attribute, true
			}
		}
	}
	return Event{}, kv.Pair{}, false
}
