package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// FinishedLotteryKeyPrefix is the prefix to retrieve all FinishedLottery
	FinishedLotteryKeyPrefix = "FinishedLottery/value/"
)

// FinishedLotteryKey returns the store key to retrieve a FinishedLottery from the index fields
func FinishedLotteryKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
