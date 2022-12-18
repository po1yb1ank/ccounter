package utils

import (
	"context"

	"github.com/po1yb1ank/ccounter/pkg/storage"
	"github.com/po1yb1ank/ccounter/pkg/watcher"
)

func NewMockStorage() *MockStorage {
	return &MockStorage{Data: make(map[string]int64)}
}

type MockStorage struct {
	Data map[string]int64
}

func (m *MockStorage) Increment(ctx context.Context, key string) (int64, error) {
	if _, ok := m.Data[key]; !ok {
		return 0, storage.ErrorKeyNotFound
	}

	m.Data[key] += 1
	return m.Data[key], nil
}

func (m *MockStorage) Decrement(ctx context.Context, key string) (int64, error) {
	if _, ok := m.Data[key]; !ok {
		return 0, storage.ErrorKeyNotFound
	}

	m.Data[key] -= 1
	return m.Data[key], nil
}

func (m *MockStorage) Set(ctx context.Context, key string, value int64) error {
	m.Data[key] = value
	return nil
}

func (m *MockStorage) Current(ctx context.Context, key string) (int64, error) {
	return m.Data[key], nil
}

type MockLogger struct{}

func (m *MockLogger) Debug(string) {}
func (m *MockLogger) Info(string)  {}
func (m *MockLogger) Error(string) {}

type MockWatcher struct{}

func (m *MockWatcher) AddSubscriber(s watcher.ISubscriber)        {}
func (m *MockWatcher) NotifyChange(key string, value int64) error { return nil }
