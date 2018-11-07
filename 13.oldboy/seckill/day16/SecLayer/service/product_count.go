package service

import (
	"sync"
)

type ProductCountMgr struct {
	productCount map[int]int //key是商品id value是此商品的销售数量
	lock         sync.RWMutex
}

func NewProductCountMgr() (productMgr *ProductCountMgr) {
	productMgr = &ProductCountMgr{
		productCount: make(map[int]int, 128),
	}

	return
}

func (p *ProductCountMgr) Count(productId int) (count int) {
	p.lock.RLock()
	defer p.lock.RUnlock()

	count = p.productCount[productId]
	return

}

func (p *ProductCountMgr) Add(productId, count int) {

	p.lock.Lock()
	defer p.lock.Unlock()

	cur, ok := p.productCount[productId]
	if !ok {
		cur = count
	} else {
		cur += count
	}

	p.productCount[productId] = count
}
