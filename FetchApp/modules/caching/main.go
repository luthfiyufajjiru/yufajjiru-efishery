package caching

import (
	"sync"
)

type SessionCaching struct {
	FetchData func() interface{}
	data      *interface{}
	mtx       *sync.RWMutex
}

func (s *SessionCaching) Begin() {
	s.mtx = &sync.RWMutex{}
}

func (s *SessionCaching) CleanCache() {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	s.data = nil
}

func (s *SessionCaching) SetData(inp interface{}) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	*s.data = inp
}

func (s *SessionCaching) GetData() (res *interface{}, err error) {
	s.mtx.RLock()
	if s.data == nil {
		s.mtx.RUnlock()
		s.mtx.Lock()

		x := s.FetchData()
		s.data = &x
		res = s.data

		s.mtx.Unlock()
	} else if s.data != nil {
		res = s.data
		s.mtx.RUnlock()
	}

	if c := s.mtx.TryRLock(); !c {
		s.mtx.RUnlock()
	}

	return
}
