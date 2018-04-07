package cxtgo

import (
	"sync"
)

// Orderbook is a definition for an orderbook
type Orderbook interface {
	Symbol() Symbol
	Head(n int)
}

type ConcurrentOrderbook struct {
	sync.RWMutex

	symbol Symbol
}

func (co *ConcurrentOrderbook) Symbol() Symbol {
	return co.symbol
}

func (co *ConcurrentOrderbook) Get() {
	panic("not implemented")
}

func (co *ConcurrentOrderbook) Head() {
	panic("not implemented")
}
