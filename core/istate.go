package core

import (
	"fmt"
	"sync"
)

var (
	internalStates sync.Map
	elemCounter    int
	elemMu         sync.Mutex
)

type StateMap struct {
	m sync.Map
}

func newStateMap() *StateMap { return &StateMap{} }

func (s *StateMap) Get(key string) (any, bool) { return s.m.Load(key) }
func (s *StateMap) Set(key string, val any)     { s.m.Store(key, val) }
func (s *StateMap) Del(key string)              { s.m.Delete(key) }
func (s *StateMap) Int(key string, def int) int {
	if v, ok := s.m.Load(key); ok {
		return v.(int)
	}
	return def
}
func (s *StateMap) Str(key string, def string) string {
	if v, ok := s.m.Load(key); ok {
		return v.(string)
	}
	return def
}
func (s *StateMap) Bool(key string, def bool) bool {
	if v, ok := s.m.Load(key); ok {
		return v.(bool)
	}
	return def
}

func (e *Element) hash() string {
	if e.hsh == "" {
		elemMu.Lock()
		elemCounter++
		e.hsh = fmt.Sprintf("el_%d", elemCounter)
		elemMu.Unlock()
	}
	return e.hsh
}

func (e *Element) InitState() *StateMap {
	h := e.hash()
	if v, ok := internalStates.Load(h); ok {
		return v.(*StateMap)
	}
	s := newStateMap()
	internalStates.Store(h, s)
	return s
}

func (e *Element) State() *StateMap { return e.InitState() }
