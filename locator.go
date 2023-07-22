package locator

import (
	"fmt"
	"sync"
)

var (
	instance *locator
	once     sync.Once
)

type locator struct {
	m        sync.RWMutex
	services map[string]any
}

func (l *locator) register(t any, s any) {
	name := l.getTypeName(t)
	l.m.Lock()
	l.services[name] = s
	l.m.Unlock()
}

func (l *locator) resolve(t any) (any, error) {
	name := l.getTypeName(t)
	l.m.RLock()
	s, ok := l.services[name]
	l.m.RUnlock()

	if !ok {
		return nil, fmt.Errorf("service not found: %T", t)
	}
	return s, nil
}

func (l *locator) getTypeName(myvar any) string {
	return fmt.Sprintf("%T", myvar)
}

func (l *locator) list() []interface{} {
	l.m.RLock()
	defer l.m.RUnlock()

	var list []interface{}
	for _, v := range l.services {
		list = append(list, v)
	}
	return list
}

func newServiceLocator() *locator {
	once.Do(func() {
		instance = &locator{
			m:        sync.RWMutex{},
			services: map[string]interface{}{},
		}
	})
	return instance
}

func Register[T any](s T) {
	l := newServiceLocator()
	var t T
	l.register(&t, s)
}

func Resolve[T any]() (T, error) {
	l := newServiceLocator()
	var t T
	obj, err := l.resolve(&t)
	if err != nil {
		return t, err
	}
	return obj.(T), err
}

func List() []any {
	l := newServiceLocator()
	return l.list()
}
