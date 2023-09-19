package myMap

import (
	"context"
	"sync"
	"time"
)

type Mychan struct {
	sync.Once

	ch chan struct{}
}

func NewMychan() *Mychan {
	return &Mychan{
		ch: make(chan struct{}),
	}
}

func (m *Mychan) Close() {
	m.Do(func() {
		close(m.ch)
	})
}

type MyConMap struct {
	sync.Mutex
	mp        map[int]int
	keyTochan map[int]*Mychan
}

func NewMyConMap() *MyConMap {
	return &MyConMap{
		mp:        make(map[int]int),
		keyTochan: make(map[int]*Mychan),
	}
}

func (m *MyConMap) Put(k, v int) {
	m.Lock()
	defer m.Unlock()
	m.mp[k] = v

	ch, ok := m.keyTochan[k]
	if !ok {
		return
	}
	ch.Close()
}

func (m *MyConMap) Get(k int, MaxWaitingDuration time.Duration) (int, error) {
	m.Lock()
	v, ok := m.mp[k]
	if ok {
		m.Unlock()
		return v, nil
	}
	ch, ok := m.keyTochan[k]
	if !ok {
		ch = NewMychan()
		m.keyTochan[k] = ch
	}
	tCx, cancle := context.WithTimeout(context.Background(), MaxWaitingDuration)
	defer cancle()

	m.Unlock()

	select {
	case <-tCx.Done():
		return -1, tCx.Err()
	case <-ch.ch:
	}

	m.Lock()
	v = m.mp[k]
	m.Unlock()
	return v, nil
}
