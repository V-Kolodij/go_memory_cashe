package gomemorycache

type Gocache struct {
	cacheItems map[string]CacheItems
}

type CacheItems struct {
	Value interface{}
}

func New() *Gocache {
	return &Gocache{
		cacheItems: make(map[string]CacheItems),
	}
}
