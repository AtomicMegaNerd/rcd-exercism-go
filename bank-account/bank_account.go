package account

import "sync"

type Account struct {
	balance int64
	open    bool
	mux     sync.RWMutex
}

func Open(amount int64) *Account {
	if amount < 0 {
		return nil
	}
	return &Account{balance: amount, open: true, mux: sync.RWMutex{}}
}

func (a *Account) Balance() (int64, bool) {
	a.mux.RLock()
	defer a.mux.RUnlock()
	return a.balance, a.open
}

func (a *Account) Deposit(amount int64) (int64, bool) {
	a.mux.Lock()
	defer a.mux.Unlock()

	if !a.open {
		return a.balance, false
	}

	if amount+a.balance < 0 {
		return a.balance, false
	}

	a.balance += amount
	return a.balance, true
}

func (a *Account) Close() (int64, bool) {
	a.mux.Lock()
	defer a.mux.Unlock()

	if !a.open {
		return a.balance, false
	}

	payout := a.balance
	a.balance = 0
	a.open = false
	return payout, true
}
