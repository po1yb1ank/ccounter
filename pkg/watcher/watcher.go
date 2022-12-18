package watcher

import (
	"encoding/json"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/po1yb1ank/ccounter/pkg/logger"
)

type ISubscriber interface {
	WriteMessage(messageType int, data []byte) error
}

type IWatcher interface {
	AddSubscriber(subscriber ISubscriber)
	NotifyChange(key string, value int64) error
}

type WSWatcher struct {
	sync.Mutex
	subscribers []ISubscriber
	logger      logger.ILogger
}

func NewWSWatcher() *WSWatcher {
	return &WSWatcher{
		subscribers: []ISubscriber{},
	}
}

func (w *WSWatcher) AddSubscriber(subscriber ISubscriber) {
	w.Mutex.Lock()
	w.subscribers = append(w.subscribers, subscriber)
	w.Mutex.Unlock()
}

func (w *WSWatcher) NotifyChange(key string, value int64) error {
	w.Mutex.Lock()
	defer w.Mutex.Unlock()

	message, err := json.Marshal(&UpdateMessage{Key: &key, Value: &value})
	if err != nil {
		return err
	}

	for _, sub := range w.subscribers {
		sub := sub
		go func() {
			err := sub.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				w.logger.Error(err.Error())
			}
		}()
	}

	return nil
}
