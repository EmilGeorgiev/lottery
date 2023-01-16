package types

const (
	// ModuleName defines the module name
	ModuleName = "lottery"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_lottery"

	EnterLotteryGas = 5

	DeadlineLayout = "2006-01-02 15:04:05.999999999 +0000 UTC"

	TxFee = 5

	// FeeCollectorName the root string for the fee collector account address
	FeeCollectorName = "fee_collector"
)

const (
	EnterLotteryEventType = "enter-lottery"
	EnterLotteryEventUser = "user"
	EnterLotteryEventBet  = "bet"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	LotteryKey = "Lottery-value-"
)

const (
	SystemInfoKey = "SystemInfo-value-"
)
