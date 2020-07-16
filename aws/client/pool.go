package awsclient

import "sync"

var pool []*Client

var mutex = &sync.Mutex{}

// FromPool provides an aws client from the pool
func FromPool() *Client {
	if len(pool) == 0 {
		return new()
	}

	var client *Client

	mutex.Lock()
	client, pool = pool[0], pool[1:] // shift
	mutex.Unlock()

	return client
}

// ToPool returns an aws client to the pool
func ToPool(client *Client) {
	mutex.Lock()
	pool = append(pool, client) // push
	mutex.Unlock()
}
