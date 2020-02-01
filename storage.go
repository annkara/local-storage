package main

import "sync"

var (
	mu sync.Mutex
	d  map[string]string
)

func init() {
	d = make(map[string]string)
}

func get(key string) string {
	mu.Lock()
	defer mu.Unlock()
	v, ok := d[key]
	if !ok {
		return ""
	}
	return v
}

func postOrPut(key string, value string) {
	mu.Lock()
	defer mu.Unlock()
	d[key] = value
}

func del(key string) {
	mu.Lock()
	defer mu.Unlock()
	delete(d, key)
}
