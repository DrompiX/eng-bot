package states

import (
	"sync"
)

type MemDB struct {
	mu        *sync.RWMutex
	userState map[int64]State
}

func NewMemDB() *MemDB {
	return &MemDB{
		mu:        new(sync.RWMutex),
		userState: make(map[int64]State),
	}
}

func (m *MemDB) GetState(chatId int64) State {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if s, exists := m.userState[chatId]; exists {
		return s
	}
	return &DefaultState{}
}

func (m *MemDB) UpdateState(chatId int64, s State) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.userState[chatId] = s
	return nil
}