package core

import "sync"

type Store struct {
	mu   sync.RWMutex
	data map[string]any
	subs map[string][]func(any)
}

var global = &Store{
	data: make(map[string]any),
	subs: make(map[string][]func(any)),
}

func Global() *Store { return global }

func Set(key string, val any)            { global.Set(key, val) }
func Get(key string) any                 { return global.Get(key) }
func Del(key string)                     { global.Del(key) }
func Int(key string, def int) int        { return global.Int(key, def) }
func Str(key string, def string) string  { return global.Str(key, def) }
func Bool(key string, def bool) bool     { return global.Bool(key, def) }
func On(key string, fn func(any))        { global.On(key, fn) }

func (s *Store) Set(key string, val any) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[key] = val
	if subs, ok := s.subs[key]; ok {
		for _, fn := range subs {
			fn(val)
		}
	}
}

func (s *Store) Get(key string) any {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.data[key]
}

func (s *Store) Del(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.data, key)
}

func (s *Store) Int(key string, def int) int {
	v := s.Get(key)
	if v == nil {
		return def
	}
	return v.(int)
}

func (s *Store) Str(key string, def string) string {
	v := s.Get(key)
	if v == nil {
		return def
	}
	return v.(string)
}

func (s *Store) Bool(key string, def bool) bool {
	v := s.Get(key)
	if v == nil {
		return def
	}
	return v.(bool)
}

func (s *Store) On(key string, fn func(any)) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.subs[key] = append(s.subs[key], fn)
}
