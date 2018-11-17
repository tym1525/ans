package session

import (
	"github.com/pborman/uuid"
	"sync"
)

type MemorySessionMgr struct {
	sessionMap map[string]Session
	rwlock     sync.RWMutex
}

func (s *MemorySessionMgr) Get(sessionId string) (session Session, err error) {
	s.rwlock.RLock()
	defer s.rwlock.RUnlock()

	session, ok := s.sessionMap[sessionId]
	if !ok {
		err = ErrSessionNotExist
		return
	}
	return
}

func (s *MemorySessionMgr) CreateSession() (session Session, err error) {
	s.rwlock.Lock()
	defer s.rwlock.Unlock()

	id := uuid.NewUUID()
	sessionId := id.String()
	session = NewMemorySession(sessionId)
	s.sessionMap[sessionId] = session
	return
}
