package sse

import (
	"sync"

	"github.com/kajtekajtek/forum/backend/internal/models"
)

type Manager struct {
	mu sync.RWMutex
	// map[channelID] = channel's subscribers message chans
	subscribers map[uint]map[chan models.Message]struct{}
}

func NewManager() *Manager {
	return &Manager{subscribers: make(map[uint]map[chan models.Message]struct{})}
}

/*
	Subscribe creates new message chan and assigns it to 
	channel's subscribers list
*/
func (m *Manager) Subscribe(channelID uint) chan models.Message {
	ch := make(chan models.Message, 1)

	m.mu.Lock()
	defer m.mu.Unlock()

	subs, ok := m.subscribers[channelID]
	if !ok {
		subs = make(map[chan models.Message]struct{})
		m.subscribers[channelID] = subs
	}
	subs[ch] = struct{}{}
	return ch
}

/*
	Unsubscribe deletes the channel's subscriber message chan and 
	eventually deletes the channel's subscribers list
*/
func (m *Manager) Unsubscribe(channelID uint, ch chan models.Message) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if subs, ok := m.subscribers[channelID]; ok {
		if _, ok := subs[ch]; ok {
			if _, ok := subs[ch]; ok {
				delete(subs, ch)
				close(ch)
			}
			if len(subs) == 0 {
				delete(m.subscribers, channelID)
			}
		}
	}
}

/*
	Publish sends the message to all of the channel's subscribers message chans
*/
func (m* Manager) Publish(channelID uint, msg models.Message) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if subs, ok := m.subscribers[channelID]; ok {
		for ch := range subs {
			select {
			case ch <- msg:
			default:
			}
		}
	}
}
