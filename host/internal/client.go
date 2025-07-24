package internal

import (
	"fmt"
	"sync"

	"github.com/owenthereal/upterm/host/api"
)

func NewClientRepo() *ClientRepo {
	return &ClientRepo{}
}

type ClientRepo struct {
	clients sync.Map
	mu      sync.Mutex
}

func (c *ClientRepo) Add(client *api.Client) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	_, loaded := c.clients.LoadOrStore(client.Id, client)
	if loaded {
		return fmt.Errorf("client already exists")
	}

	return nil
}

func (c *ClientRepo) Delete(clientId string) {
	c.clients.Delete(clientId)
}

func (c *ClientRepo) Get(clientId string) *api.Client {
	val, _ := c.clients.Load(clientId)
	if val != nil {
		return val.(*api.Client)
	}

	return nil
}

func (c *ClientRepo) Clients() []*api.Client {
	var clients []*api.Client

	c.clients.Range(func(key, value interface{}) bool {
		cc := value.(*api.Client)
		clients = append(clients, cc)
		return true
	})

	return clients
}
