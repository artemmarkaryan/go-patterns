package proxy

import "time"

type DB interface {
	GetData() string
}

type ProdDB struct {
	data string
}

func (p ProdDB) GetData() string {
	return p.data
}

type ProdDBProxy struct {
	db ProdDB
	cachedAt time.Time
	cache string
}

func (p ProdDBProxy) GetData() string {
	if time.Now().Sub(p.cachedAt) > time.Second {
		p.cache = p.db.GetData()
		p.cachedAt = time.Now()
	}
	return p.cache
}
