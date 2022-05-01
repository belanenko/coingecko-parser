package model

type PricePerTimestamp struct {
	Timestamp int64
	Price     float32
}

type History struct {
	Wallets map[string][]PricePerTimestamp
}
