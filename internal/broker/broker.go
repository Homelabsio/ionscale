package broker

import (
	"sync"
	"tailscale.com/tailcfg"
	"tailscale.com/types/key"
)

type BrokerPool struct {
	lock  sync.Mutex
	store map[uint64]Broker
}

type Signal struct {
	PeerUpdated  *uint64
	PeersRemoved []uint64
	ACLUpdated   bool
	DNSUpdated   bool
	derpMap      *tailcfg.DERPMap
}

type Broker interface {
	AddClient(*Client)
	RemoveClient(uint64)

	SignalUpdate()
	SignalPeerUpdated(id uint64)
	SignalPeersRemoved([]uint64)
	SignalDNSUpdated()
	SignalACLUpdated()
	SignalDERPMapUpdated(c *tailcfg.DERPMap)

	IsConnected(uint64) bool
}

func NewBrokerPool() *BrokerPool {
	return &BrokerPool{
		store: make(map[uint64]Broker),
	}
}

func (m *BrokerPool) Get(tailnetID uint64) Broker {
	m.lock.Lock()
	defer m.lock.Unlock()
	b, ok := m.store[tailnetID]
	if !ok {
		b = newBroker(tailnetID)
		m.store[tailnetID] = b
	}
	return b
}

func (m *BrokerPool) SignalDERPMapUpdated(c *tailcfg.DERPMap) {
	m.lock.Lock()
	defer m.lock.Unlock()
	for _, b := range m.store {
		b.SignalDERPMapUpdated(c)
	}
}

func (m *BrokerPool) SignalUpdate() {
	m.lock.Lock()
	defer m.lock.Unlock()
	for _, b := range m.store {
		b.SignalUpdate()
	}
}

func newBroker(tailnetID uint64) Broker {
	b := &broker{
		tailnetID:      tailnetID,
		newClients:     make(chan *Client),
		closingClients: make(chan uint64),
		clients:        make(map[uint64]*Client),
		signalChannel:  make(chan *Signal),
	}

	go b.listen()

	return b
}

type broker struct {
	tailnetID      uint64
	privateKey     *key.MachinePrivate
	newClients     chan *Client
	closingClients chan uint64
	signalChannel  chan *Signal
	clients        map[uint64]*Client
}

func (h *broker) IsConnected(id uint64) (ok bool) {
	_, ok = h.clients[id]
	return
}

func (h *broker) AddClient(client *Client) {
	h.newClients <- client
}

func (h *broker) RemoveClient(id uint64) {
	h.closingClients <- id
}

func (h *broker) SignalUpdate() {
	h.signalChannel <- &Signal{}
}

func (h *broker) SignalPeerUpdated(id uint64) {
	h.signalChannel <- &Signal{PeerUpdated: &id}
}

func (h *broker) SignalPeersRemoved(ids []uint64) {
	h.signalChannel <- &Signal{PeersRemoved: ids}
}

func (h *broker) SignalDNSUpdated() {
	h.signalChannel <- &Signal{DNSUpdated: true}
}

func (h *broker) SignalACLUpdated() {
	h.signalChannel <- &Signal{ACLUpdated: true}
}

func (h *broker) SignalDERPMapUpdated(c *tailcfg.DERPMap) {
	h.signalChannel <- &Signal{derpMap: c}
}

func (h *broker) listen() {
	for {
		select {
		case s := <-h.newClients:
			h.clients[s.id] = s
		case s := <-h.closingClients:
			delete(h.clients, s)
		case s := <-h.signalChannel:
			for _, c := range h.clients {
				c.SignalUpdate(s)
			}
		}
	}
}
