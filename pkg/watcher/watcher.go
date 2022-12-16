package watcher

import (
	"encoding/json"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/po1yb1ank/ccounter/pkg/logger"
)

type Watcher struct {
	sync.Mutex
	subscribers []*websocket.Conn
	logger      logger.ILogger
}

func NewWatcher() *Watcher {
	return &Watcher{
		subscribers: []*websocket.Conn{},
	}
}

func (w *Watcher) AddSubscriber(subscriber *websocket.Conn) {
	w.Mutex.Lock()
	w.subscribers = append(w.subscribers, subscriber)
	w.Mutex.Unlock()
}

func (w *Watcher) NotifyChange(key string, value int64) error {
	w.Mutex.Lock()
	defer w.Mutex.Unlock()

	message, _ := json.Marshal(&UpdateMessage{Key: key, Value: value})

	for _, sub := range w.subscribers {
		err := sub.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			w.logger.Error(err.Error())
		}
	}

	return nil
}
