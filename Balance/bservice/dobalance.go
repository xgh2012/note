package bservice

import (
	"sync"
)

type Balances struct {
	AppId string
}

var (
	AppIdBalance = make(map[string]int)          //商户余额
	AppIdReadMu  = make(map[string]sync.RWMutex) //商户读锁
	AppIdWriteMu = make(map[string]sync.Mutex)   //商户写锁
)

//读取商户余额
func (b *Balances) GetBalance() int {
	c := AppIdReadMu[b.AppId]
	c.RLock()
	balance := AppIdBalance[b.AppId]
	c.RUnlock()
	return balance
}

//更新余额
func (b *Balances) AddBalance(add int) int {
	c := AppIdWriteMu[b.AppId]
	c.Lock()
	AppIdBalance[b.AppId] += add
	balance := AppIdBalance[b.AppId]
	c.Unlock()
	return balance
}
