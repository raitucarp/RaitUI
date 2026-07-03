package core

import (
	"fmt"
	"sync"
)

var (
	globalStore sync.Map
	localStore  sync.Map
)

func SetState(key string, value any) {
	globalStore.Store(key, value)
}

func GetState(key string) (any, bool) {
	return globalStore.Load(key)
}

func DeleteState(key string) {
	globalStore.Delete(key)
}

func SetLocalState(componentID, key string, value any) {
	localStore.Store(localKey(componentID, key), value)
}

func GetLocalState(componentID, key string) (any, bool) {
	return localStore.Load(localKey(componentID, key))
}

func DeleteLocalState(componentID, key string) {
	localStore.Delete(localKey(componentID, key))
}

func localKey(componentID, key string) string {
	return fmt.Sprintf("component_%s_%s", componentID, key)
}

func StateInt(key string, def int) func() int {
	return StateGetter(key, def)
}

func StateStr(key string, def string) func() string {
	return StateGetter(key, def)
}

func Setter[T any](key string) func(T) {
	return func(v T) {
		SetState(key, v)
	}
}

func StateGetter[T any](key string, def T) func() T {
	if _, ok := globalStore.Load(key); !ok {
		globalStore.Store(key, def)
	}
	return func() T {
		v, ok := globalStore.Load(key)
		if !ok {
			return def
		}
		return v.(T)
	}
}

func LocalStateInt(componentID, key string, def int) func() int {
	return LocalStateGetter(componentID, key, def)
}

func LocalStateStr(componentID, key string, def string) func() string {
	return LocalStateGetter(componentID, key, def)
}

func LocalStateGetter[T any](componentID, key string, def T) func() T {
	lk := localKey(componentID, key)
	if _, ok := localStore.Load(lk); !ok {
		localStore.Store(lk, def)
	}
	return func() T {
		v, ok := localStore.Load(lk)
		if !ok {
			return def
		}
		return v.(T)
	}
}

func LocalSetter[T any](componentID, key string) func(T) {
	return func(v T) {
		SetLocalState(componentID, key, v)
	}
}
