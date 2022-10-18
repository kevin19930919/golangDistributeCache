package geecache

import "sync"

type Group struct {
	name      string
	getter    Getter
	mainCache cache
}

var mu sync.RWMutex
var groups = make(map[string]*Group)

type Getter interface {
	Get(key string) ([]byte, error)
}

func NewGroup(name string, cacheBytes int64, getter Getter) *Group {
	if getter == nil {
		panic("nil getter")
	}
	mu.Lock()
	defer mu.Unlock()
	g := &Group{
		name:      name,
		getter:    getter,
		mainCache: cache{cacheBytes: cacheBytes},
	}
	groups[name] = g
	return g
}

// cause we hope to have a multiple read but only one write condition, so use RWMutex instead of Mutex
func GetGroup(name string) *Group {
	mu.RLock()
	g := groups[name]
	mu.RUnlock()
	return g
}

type GetterFunc func(key string) ([]byte, error)

func (f GetterFunc) Get(key string) ([]byte, error) {
	return f(key)
}
