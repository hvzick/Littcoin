package types

import "time"

type TransactionManagerChainConfig interface {
	BroadcasterChainConfig
}

type TransactionManagerFeeConfig interface {
	BroadcasterFeeConfig
	ConfirmerFeeConfig
}

type TransactionManagerTransactionsConfig interface {
	BroadcasterTransactionsConfig
	ConfirmerTransactionsConfig
	ResenderTransactionsConfig
	ReaperTransactionsConfig

	ForwardersEnabled() bool
	MaxQueued() uint64
}

type BroadcasterChainConfig interface {
	IsL2() bool
}

type BroadcasterFeeConfig interface {
	MaxFeePrice() string     // logging value
	FeePriceDefault() string // logging value
}

type BroadcasterTransactionsConfig interface {
	MaxInFlight() uint32
}

type BroadcasterListenerConfig interface {
	FallbackPollInterval() time.Duration
}

type ConfirmerFeeConfig interface {
	BumpTxDepth() uint32
	LimitDefault() uint64

	// from gas.Config
	BumpThreshold() uint64
	MaxFeePrice() string // logging value
}

type ConfirmerDatabaseConfig interface {
	// from pg.QConfig
	DefaultQueryTimeout() time.Duration
}

type ConfirmerTransactionsConfig interface {
	MaxInFlight() uint32
	ForwardersEnabled() bool
}

type ResenderChainConfig interface {
	RPCDefaultBatchSize() uint32
}

type ResenderTransactionsConfig interface {
	ResendAfterThreshold() time.Duration
	MaxInFlight() uint32
}

type ReaperTransactionsConfig interface {
	ReaperInterval() time.Duration
	ReaperThreshold() time.Duration
}
